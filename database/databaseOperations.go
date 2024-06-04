package database

import (
	"dashboard/utils"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database interface {
	CreateUser(user utils.UserRegistration) error
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

func (db *MongoDB) CreateUser(user utils.UserRegistration) error {
	return nil
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
func (m *MockDB) CreateUser(user utils.UserRegistration) error {
	m.users[user.Username] = user
	return nil
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
