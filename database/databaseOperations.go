package database

import (
	"context"
	"dashboard/utils"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"strings"
)

type Database interface {

	// Functions related to user registration
	CreateUser(user utils.UserRegistration) (int, error)
	ReadUser(username string) (utils.UserRegistration, error)
	UpdateUser(user utils.UserRegistration) error
	DeleteUser(username string) error
}

// MongoDB is a struct for the actual MongoDB database
type MongoDB struct {
	Client *mongo.Client
	dbName string
	users  *mongo.Collection
}

// NewMongoDB instantiates a new MongoDB
func NewMongoDB(client *mongo.Client, dbName string, collection string) *MongoDB {
	return &MongoDB{
		Client: client,
		dbName: dbName,
		users:  client.Database(dbName).Collection(collection),
	}
}

// CreateUser creates a new user in the database
func (db *MongoDB) CreateUser(user utils.UserRegistration) (int, error) {

	// Enforce required fields
	if user.Username == "" || user.Password == "" || user.Email == "" {
		log.Println("username, password, and email are required fields")
		return http.StatusBadRequest, errors.New("username, password, and email are required fields")
	}

	// Set username and email to lowercase
	user.Username = strings.ToLower(user.Username)
	user.Email = strings.ToLower(user.Email)

	// Check if the username or email already exists
	if utils.CheckUsernameAndEmail(user) {
		log.Println("username or email already exists")
		return http.StatusBadRequest, errors.New("username or email already exists")
	}

	// Enforce a password only containing characters '1234567890'
	if !utils.EnforcePassword(user.Password) {
		log.Println("please don't use an actual password for this. The only accepted characters are '1234567890'")
		return http.StatusBadRequest, errors.New("please don't use an actual password for this. The only accepted characters are '1234567890'")
	}

	var err error

	// Hash the password
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		log.Println("error hashing password")
		return http.StatusInternalServerError, errors.New("error hashing password")
	}

	// Open the collection and insert the user
	collection := db.Client.Database(db.dbName).Collection(utils.COLLECTION_USERS)
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	log.Println("User '" + user.Username + "' registered.")

	return http.StatusCreated, nil
}

func (db *MongoDB) ReadUser(username string) (utils.UserRegistration, error) {
	return utils.UserRegistration{}, nil
}

func (db *MongoDB) UpdateUser(user utils.UserRegistration) error {
	return nil
}

func (db *MongoDB) DeleteUser(username string) error {
	return nil
}

// MockDB is a database struct for testing
type MockDB struct {
	users map[string]utils.UserRegistration
}

// NewMockDB instantiates a new MockDB
func NewMockDB() *MockDB {
	return &MockDB{
		users: make(map[string]utils.UserRegistration),
	}
}

// CreateUser creates a new user in the database
func (m *MockDB) CreateUser(user utils.UserRegistration) (int, error) {

	// Enforce required fields
	if user.Username == "" || user.Password == "" || user.Email == "" {
		log.Println("username, password, and email are required fields")
		return http.StatusBadRequest, errors.New("username, password, and email are required fields")
	}

	// Set username and email to lowercase
	user.Username = strings.ToLower(user.Username)
	user.Email = strings.ToLower(user.Email)

	// Check if the username or email already exists
	if _, exists := m.users[user.Username]; exists {
		log.Println("username or email already exists")
		return http.StatusBadRequest, errors.New("username or email already exists")
	}

	// Enforce a password only containing characters '1234567890'
	if !utils.EnforcePassword(user.Password) {
		log.Println("please don't use an actual password for this. The only accepted characters are '1234567890'")
		return http.StatusBadRequest, errors.New("please don't use an actual password for this. The only accepted characters are '1234567890'")
	}

	var err error

	// Hash the password
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		log.Println("error hashing password")
		return http.StatusInternalServerError, errors.New("error hashing password")
	}

	// Insert the user into the test database
	m.users[user.Username] = user

	log.Println("User '" + user.Username + "' registered.")

	return http.StatusCreated, nil
}

// ReadUser reads a user from the database
func (m *MockDB) ReadUser(username string) (utils.UserRegistration, error) {
	user, exists := m.users[username]
	if !exists {
		return utils.UserRegistration{}, fmt.Errorf("user %s does not exist", username)
	}
	return user, nil
}

// UpdateUser updates a user in the database
func (m *MockDB) UpdateUser(user utils.UserRegistration) error {
	_, exists := m.users[user.Username]
	if !exists {
		return fmt.Errorf("user %s does not exist", user.Username)
	}

	m.users[user.Username] = user
	return nil
}

// DeleteUser deletes a user from the database
func (m *MockDB) DeleteUser(username string) error {
	_, exists := m.users[username]
	if !exists {
		return fmt.Errorf("user %s does not exist", username)
	}
	delete(m.users, username)
	return nil
}
