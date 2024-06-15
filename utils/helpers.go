package utils

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"strings"
)

var Client *mongo.Client

// DBConnect Function to connect to the MongoDB database
func DBConnect() *mongo.Client {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoURI := os.Getenv("MONGODB_URI")
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	var err error
	// Create a new client and connect to the server
	Client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// Get a handle to the users connection
	collection := Client.Database(COLLECTION_USERS).Collection(COLLECTION_USERS)

	// Establish uniqueness constraint on username
	usernameIndex := mongo.IndexModel{
		Keys:    bson.M{"username": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err = collection.Indexes().CreateOne(context.TODO(), usernameIndex)
	if err != nil {
		log.Fatal(err)
	}
	// Establish uniqueness constraint on email
	emailIndex := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err = collection.Indexes().CreateOne(context.TODO(), emailIndex)
	if err != nil {
		log.Fatal(err)
	}

	// Send a ping to confirm a successful connection
	if err := Client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("\nService successfully connected to MongoDB.\n")

	return Client
}

// EnforcePassword Function to ensure users don't use an actual password
func EnforcePassword(password string) bool {

	accepted := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

	// Ensure that all characters in the password are of the accepted characters (numbers)
	for _, char := range password {
		found := false
		for _, element := range accepted {
			if string(char) == element {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// HashPassword Function to hash a password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPassword Function to check if a password matches a hash
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// EnsureCorrectPath appends a '/' to the path if it is missing
func EnsureCorrectPath(r *http.Request) {
	if r.URL.Path[len(r.URL.Path)-1] != '/' {
		r.URL.Path += "/"
	}
}

// CheckUsernameAndEmail Function to check if a username or email already exists
func CheckUsernameAndEmail(user UserRegistration) bool {

	collection := Client.Database(COLLECTION_USERS).Collection(COLLECTION_USERS)
	existingUser := UserRegistration{}
	err := collection.FindOne(context.TODO(), bson.M{"$or": []bson.M{
		{"username": user.Username},
		{"email": user.Email},
	}}).Decode(&existingUser)

	if err != nil {
		return false
	}
	return true
}

// CheckEmail Function to check if an email already exists
func CheckEmail(user UserRegistration) bool {

	collection := Client.Database(COLLECTION_USERS).Collection(COLLECTION_USERS)
	existingUser := UserRegistration{}
	err := collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existingUser)

	if err != nil {
		return false
	}
	return true

}

// ExtractUsername Function to extract the username from the URL path
func ExtractUsername(w http.ResponseWriter, r *http.Request) string {
	// Extract the username from the URL path
	path := strings.Split(r.URL.Path, "/")
	username := path[len(path)-2]

	// Enforce username to be specified
	if username == ENDPOINT_REGISTRATIONS || username == ENDPOINT_SET_COOKIE || username == ENDPOINT_GET_COOKIE || username == ENDPOINT_DELETE_COOKIE || username == ENDPOINT_WEATHER {
		http.Error(w, "Username must be provided", http.StatusBadRequest)
		log.Println("Username not provided")
		username = ""
	}

	return strings.ToLower(username)
}

// SetToLower Function to set the username and email to lowercase (using pointer and reference)
func SetToLower(user *UserRegistration) {
	user.Username = strings.ToLower(user.Username)
	user.Email = strings.ToLower(user.Email)
}

// GenerateRandomToken generates a random string of the given length
func GenerateRandomToken(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

/*
ParseFile Reads a given file and returns content as byte array.
------------------------------------------------------------------
This function was originally written by Author Christopher Frantz
for the PROG2005 course at NTNU, Gjøvik
*/
func ParseFile(filename string) []byte {
	file, e := os.ReadFile(filename)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	return file
}

// SetHeaders Function to set the headers for the response
func SetHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}

// GetCity function finds the city in the country we are looking for
func GetCity(cities []GeoCodeResponse, countryCode string) (GeoCodeResponse, bool) {
	for _, city := range cities {
		if city.CountryCode == countryCode {
			return city, true
		}
	}
	return GeoCodeResponse{}, false
}

// GetCountry finds the location of the country
func GetCountry(country string) (GeoCodeResponse, error) {

	// Fetch the API response from the country
	geoGet, err := http.Get(GEOCODING_API + country)
	if err != nil {
		return GeoCodeResponse{}, err
	}
	defer geoGet.Body.Close().Error()

	// Decode the response
	var geoCodeResponse []GeoCodeResponse
	err = json.NewDecoder(geoGet.Body).Decode(&geoCodeResponse)
	if err != nil {
		return GeoCodeResponse{}, err
	}

	// Return the first element in the response
	return geoCodeResponse[0], nil
}
