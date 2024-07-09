package utils

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
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

	// The services endpoints in a slice
	var endpoints = []string{
		ENDPOINT_REGISTRATIONS,
		ENDPOINT_SET_COOKIE,
		ENDPOINT_GET_COOKIE,
		ENDPOINT_DELETE_COOKIE,
		ENDPOINT_WEATHER,
	}

	// Enforce username to be specified
	if stringInSlice(username, endpoints) {
		http.Error(w, "Username must be provided", http.StatusBadRequest)
		log.Println("Username not provided")
		username = ""
	}

	return strings.ToLower(username)
}

// stringInSlice Function to check if a string is in a slice
func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
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
func GetCity(cities GeoCodeResults, countryCode string) (GeoCodeResponse, bool) {
	for _, city := range cities.Results {
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

	// Decode the response
	var geoCodeResponse GeoCodeResponse
	err = json.NewDecoder(geoGet.Body).Decode(&geoCodeResponse)
	if err != nil {
		return GeoCodeResponse{}, err
	}

	// Return the first element in the response
	return geoCodeResponse, nil
}

// GenerateWeatherURL generates the URL for the weather API
func GenerateWeatherURL(coordinates Coordinates) string {

	latitude := strconv.FormatFloat(coordinates.Latitude, 'f', -1, 64)
	longitude := strconv.FormatFloat(coordinates.Longitude, 'f', -1, 64)

	Url := WEATHER_API_BASE + WEATHER_API_LAT + latitude + WEATHER_API_LON + longitude + WEATHER_API_PARAMETERS

	return Url
}

// SetCurrentWeather sets the current weather
func SetCurrentWeather(weather WeatherData) DailyWeather {
	t := time.Now()
	hour := 0

	// Find the current hour
	for hour < 24 {
		if ExtractHour(weather.Hourly.Time[hour]) == stringifyHour(t.Hour()) {
			break
		}
		hour++
	}

	dailyWeather := DailyWeather{}
	dailyWeather.Date = ExtractDate(weather.Hourly.Time[hour])

	// Set the weather for the remainder of the day
	for hour < 24 {
		hourlyWeather := HourlyWeather{
			Hour:          ExtractHour(weather.Hourly.Time[hour]),
			Temperature:   weather.Hourly.Temperature[hour],
			Precipitation: weather.Hourly.Precipitation[hour],
			CloudCover:    weather.Hourly.CloudCover[hour],
			WindSpeed:     weather.Hourly.WindSpeed[hour],
		}
		hourlyWeather.Condition = DetermineWeatherCondition(hourlyWeather)
		dailyWeather.Hours = append(dailyWeather.Hours, hourlyWeather)
		hour++
	}

	return dailyWeather
}

// stringifyHour prepends a 0 to single digit hours
func stringifyHour(hour int) string {
	if hour < 10 {
		return "0" + strconv.Itoa(hour)
	}
	return strconv.Itoa(hour)
}

// ExtractHour extracts the hour from the date string
func ExtractHour(date string) string {
	t := StringToTime(date)

 	// Make sure single digits hours have a prepended 0
	return stringifyHour(t.Hour())
}

// ExtractDate extracts the date from the date string
func ExtractDate(date string) string {
	t := StringToTime(date)
	return t.String()[:10]
}

// StringToTime converts a date string to a time.Time object
func StringToTime(date string) time.Time {

	t, err := time.Parse(TIME_FORMAT, date)
	if err != nil {
		fmt.Println("Error:", err)
		return time.Time{}
	}

	return t
}

// DetermineWeatherCondition determines the weather condition based on the hourly data
func DetermineWeatherCondition(weather HourlyWeather) string {

	// If any rain is detected, return "rainy"
	if weather.Precipitation > 0 {
		return CONDITION_RAINY
	}
	// Switch to determine cloud cover
	switch {
	case weather.CloudCover > 87:
		return CONDITION_CLOUDY
	case weather.CloudCover > 70:
		return CONDITION_MOSTLY_CLOUDY
	case weather.CloudCover > 50:
		return CONDITION_PARTLY_CLOUDY
	case weather.CloudCover > 25:
		return CONDITION_MOSTLY_SUNNY
	case weather.CloudCover > 5:
		return CONDITION_MOSTLY_CLEAR
	default:
		return CONDITION_CLEAR_DAY
	}
}
