package database

import (
	"context"
	"dashboard/utils"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"strings"
)

type Database interface {

	// Functions related to user registration
	CreateUser(user utils.UserRegistration) (int, error)
	ReadUser(username string) (int, utils.UserRegistration, error)
	UpdateUser(username string, user utils.UserRegistration) (int, error)
	DeleteUser(username string) (int, error)
	CheckUserExistence(username string) (bool, utils.UserRegistration)
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

// CreateUser creates a new user in the MongoDB database
func (db *MongoDB) CreateUser(user utils.UserRegistration) (int, error) {

	// Enforce required fields
	if user.Username == "" || user.Password == "" || user.Email == "" {
		log.Println("username, password, and email are required fields")
		return http.StatusBadRequest, errors.New("username, password, and email are required fields")
	}

	// Set username and email to lowercase
	utils.SetToLower(&user)

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

// ReadUser reads a user from the MongoDB database
func (db *MongoDB) ReadUser(username string) (int, utils.UserRegistration, error) {

	// Check if the user exists
	found, response := db.CheckUserExistence(username)
	if !found {
		log.Println("User not found")
		return http.StatusNotFound, utils.UserRegistration{}, errors.New("user not found")
	}

	return http.StatusOK, response, nil
}

// UpdateUser updates a user in the MongoDB database
func (db *MongoDB) UpdateUser(username string, user utils.UserRegistration) (int, error) {

	// Enforce username and email to be lowercase
	utils.SetToLower(&user)

	// Check if user exists
	found, _ := db.CheckUserExistence(username)
	if !found {
		log.Println("User not found")
		return http.StatusBadRequest, errors.New("user not found")
	}

	// Fetch the user from the database
	collection := db.Client.Database(utils.COLLECTION_USERS).Collection(utils.COLLECTION_USERS)
	currentValue := utils.UserRegistration{}
	err := collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&currentValue)
	if err != nil {
		log.Println("Error fetching user in PUT request")
		return http.StatusInternalServerError, errors.New("error fetching user")
	}

	log.Println("User found on username: ", currentValue.Username)

	// Disallow any attempted change of username
	if user.Username != currentValue.Username {
		log.Println("Attempted change of username")
		return http.StatusBadRequest, errors.New("username cannot be changed")
	}

	// Check if the password is changed
	if user.Password != "" || user.Password != currentValue.Password {

		// Apply constraints and hash the password if it is changed
		if !utils.EnforcePassword(user.Password) {
			log.Println("invalid password characters")
			return http.StatusBadRequest, errors.New("please don't use an actual password for this. The only accepted characters are '1234567890'")
		}

		if len(user.Password) < 8 {
			log.Println("password too short")
			return http.StatusBadRequest, errors.New("password must be at least 8 characters long")
		}

		// Hash the password
		user.Password, err = utils.HashPassword(user.Password)
		if err != nil {
			log.Println("error hashing password")
			return http.StatusInternalServerError, errors.New("error hashing password")
		}
	}

	// Update the user in the database
	_, err = collection.UpdateOne(context.TODO(), bson.M{"username": username}, bson.M{"$set": user})
	if err != nil {
		log.Println("error updating user")
		return http.StatusInternalServerError, errors.New("error updating user")
	}

	return http.StatusOK, nil
}

// DeleteUser deletes a user from the MongoDB database
func (db *MongoDB) DeleteUser(username string) (int, error) {

	// Check if user exists
	found, _ := db.CheckUserExistence(username)
	if !found {
		log.Println("User not found")
		return http.StatusNotFound, errors.New("user not found")
	}

	// Open the collection and delete the user
	collection := db.Client.Database(utils.COLLECTION_USERS).Collection(utils.COLLECTION_USERS)
	_, err := collection.DeleteOne(context.TODO(), bson.M{"username": username})
	if err != nil {
		log.Println("error deleting user")
		return http.StatusInternalServerError, errors.New("error deleting user")
	}

	log.Println("User '" + username + "' deleted.")
	return http.StatusNoContent, nil
}

// CheckUserExistence checks if a user exists in the database
func (db *MongoDB) CheckUserExistence(username string) (bool, utils.UserRegistration) {
	collection := utils.Client.Database(utils.COLLECTION_USERS).Collection(utils.COLLECTION_USERS)
	response := utils.UserRegistration{}
	err := collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&response)

	// Return false if it can't find the user, true otherwise
	if err != nil {
		return false, response
	}
	return true, response
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
func (m *MockDB) ReadUser(username string) (int, utils.UserRegistration, error) {
	user, exists := m.users[username]
	if !exists {
		return http.StatusNotFound, utils.UserRegistration{}, fmt.Errorf("user %s does not exist", username)
	}
	return http.StatusOK, user, nil
}

// UpdateUser updates a user in the database
func (m *MockDB) UpdateUser(username string, user utils.UserRegistration) (int, error) {

	// Enforce username and email to be lowercase
	utils.SetToLower(&user)

	currentValue, exists := m.users[username]
	if !exists {
		return http.StatusNotFound, fmt.Errorf("user %s does not exist", user.Username)
	}

	// Disallow any attempted change of username
	if user.Username != currentValue.Username {
		return http.StatusBadRequest, errors.New("username cannot be changed")
	}

	// Check if the password is changed
	if user.Password != "" || user.Password != currentValue.Password {

		// Apply constraints and hash the password if it is changed
		if !utils.EnforcePassword(user.Password) {
			return http.StatusBadRequest, errors.New("please don't use an actual password for this. The only accepted characters are '1234567890'")
		}

		if len(user.Password) < 8 {
			return http.StatusBadRequest, errors.New("password must be at least 8 characters long")
		}

		var err error
		// Hash the password
		user.Password, err = utils.HashPassword(user.Password)
		if err != nil {
			return http.StatusInternalServerError, errors.New("error hashing password")
		}

	}

	// Update the user in the database
	m.users[user.Username] = user
	return http.StatusOK, nil
}

// DeleteUser deletes a user from the database
func (m *MockDB) DeleteUser(username string) (int, error) {
	_, exists := m.users[username]
	if !exists {
		return http.StatusNotFound, fmt.Errorf("user %s does not exist", username)
	}
	delete(m.users, username)
	return http.StatusNoContent, nil
}

// CheckUserExistence checks if a user exists in the database
func (m *MockDB) CheckUserExistence(username string) (bool, utils.UserRegistration) {
	user, exists := m.users[username]
	return exists, user
}