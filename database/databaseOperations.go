package database

import (
	"context"
	"dashboard/utils"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Database interface {

	// Functions related to user registration
	CreateUser(user utils.UserRegistration) (int, error)
	ReadUser(username string) (int, utils.UserRegistration, error)
	UpdateUser(username string, user utils.UserRegistration) (int, error)
	DeleteUser(username string) (int, error)
	CheckUserExistence(username string) (bool, utils.UserRegistration)

	// Functions related to user authentication
	AuthenticateUser(database Database, userRequest utils.UserAuthentication) (int, error)

	// Functions related to API fetches
	GetGeoCode(country utils.Country, city string) (int, utils.Coordinates, error)
	GetWeather(country utils.Country, city string) (int, utils.WeatherData, error)
	SetWeeklyWeather(weatherData utils.WeatherData) (utils.WeeklyWeather, error)
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
		log.Println(utils.ERROR_FIELDS_REQUIRED)
		return http.StatusBadRequest, errors.New(utils.ERROR_FIELDS_REQUIRED)
	}

	// Set username and email to lowercase
	utils.SetToLower(&user)

	// Check if the username or email already exists
	if utils.CheckUsernameAndEmail(user) {
		log.Println(utils.ERROR_EXISTS)
		return http.StatusBadRequest, errors.New(utils.ERROR_EXISTS)
	}

	// Enforce a password only containing characters '1234567890'
	if !utils.EnforcePassword(user.Password) {
		log.Println(utils.ERROR_PASSWORD_INVALID)
		return http.StatusBadRequest, errors.New(utils.ERROR_PASSWORD_INVALID)
	}

	var err error

	// Hash the password
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		log.Println(utils.ERROR_PASSWORD_HASH)
		return http.StatusInternalServerError, errors.New(utils.ERROR_PASSWORD_HASH)
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
		log.Println(utils.ERROR_USER_NOT_FOUND)
		return http.StatusNotFound, utils.UserRegistration{}, errors.New(utils.ERROR_USER_NOT_FOUND)
	}

	return http.StatusOK, response, nil
}

// UpdateUser updates a user in the MongoDB database
func (db *MongoDB) UpdateUser(username string, user utils.UserRegistration) (int, error) {

	// Enforce username and email to be lowercase
	utils.SetToLower(&user)

	// Check if user exists, if not, then we assume a change of username was attempted
	found, _ := db.CheckUserExistence(username)
	if !found {
		log.Println(utils.ERROR_USER_NOT_FOUND)
		return http.StatusBadRequest, errors.New(utils.ERROR_USER_NOT_FOUND)
	}

	// Fetch the user from the database
	collection := db.Client.Database(utils.COLLECTION_USERS).Collection(utils.COLLECTION_USERS)
	currentValue := utils.UserRegistration{}
	err := collection.FindOne(context.TODO(), bson.M{utils.USERNAME: username}).Decode(&currentValue)
	if err != nil {
		log.Println(utils.ERROR_USER_FETCH)
		return http.StatusInternalServerError, errors.New(utils.ERROR_USER_FETCH)
	}

	// Check if email is changed, if it is, check if it already exists
	if user.Email != currentValue.Email {
		if utils.CheckEmail(user) {
			log.Println(utils.ERROR_EMAIL_EXISTS)
			return http.StatusBadRequest, errors.New(utils.ERROR_EMAIL_EXISTS)
		}
	}

	// Disallow any attempted change of username
	if user.Username != currentValue.Username {
		log.Println(utils.ERROR_USERNAME_CHANGE)
		return http.StatusBadRequest, errors.New(utils.ERROR_USERNAME_CHANGE)
	}

	// If the password value is "-" it will be ignored
	if user.Password != "-" {
		// Check if the password is changed
		if user.Password != "" || user.Password != currentValue.Password {

			// Apply constraints and hash the password if it is changed
			if !utils.EnforcePassword(user.Password) {
				log.Println(utils.LOG_PASSWORD_INVALID)
				return http.StatusBadRequest, errors.New(utils.ERROR_PASSWORD_INVALID)
			}

			if len(user.Password) < 8 {
				log.Println(utils.ERROR_PASSWORD_LENGTH)
				return http.StatusBadRequest, errors.New(utils.ERROR_PASSWORD_LENGTH)
			}

			// Hash the password
			user.Password, err = utils.HashPassword(user.Password)
			if err != nil {
				log.Println(utils.ERROR_PASSWORD_HASH)
				return http.StatusInternalServerError, errors.New(utils.ERROR_PASSWORD_HASH)
			}
		}
	} else {
		user.Password = currentValue.Password
	}

	// Update the user in the database
	_, err = collection.UpdateOne(context.TODO(), bson.M{utils.USERNAME: username}, bson.M{utils.BSON_SET: user})
	if err != nil {
		log.Println(utils.ERROR_USER_UPDATE)
		return http.StatusInternalServerError, errors.New(utils.ERROR_USER_UPDATE)
	}

	return http.StatusOK, nil
}

// DeleteUser deletes a user from the MongoDB database
func (db *MongoDB) DeleteUser(username string) (int, error) {

	// Check if user exists
	found, _ := db.CheckUserExistence(username)
	if !found {
		log.Println(utils.ERROR_USER_NOT_FOUND)
		return http.StatusNotFound, errors.New(utils.ERROR_USER_NOT_FOUND)
	}

	// Open the collection and delete the user
	collection := db.Client.Database(utils.COLLECTION_USERS).Collection(utils.COLLECTION_USERS)
	_, err := collection.DeleteOne(context.TODO(), bson.M{utils.USERNAME: username})
	if err != nil {
		log.Println(utils.ERROR_USER_DELETE)
		return http.StatusInternalServerError, errors.New(utils.ERROR_USER_DELETE)
	}

	log.Println("User '" + username + "' deleted.")
	return http.StatusNoContent, nil
}

// CheckUserExistence checks if a user exists in the database
func (db *MongoDB) CheckUserExistence(username string) (bool, utils.UserRegistration) {
	collection := utils.Client.Database(utils.COLLECTION_USERS).Collection(utils.COLLECTION_USERS)
	response := utils.UserRegistration{}
	err := collection.FindOne(context.TODO(), bson.M{utils.USERNAME: username}).Decode(&response)

	// Return false if it can't find the user, true otherwise
	if err != nil {
		return false, response
	}
	return true, response
}

// AuthenticateUser authenticates a user in the MongoDB database
func (db *MongoDB) AuthenticateUser(database Database, userRequest utils.UserAuthentication) (int, error) {

	// Check required fields
	if userRequest.Username == "" || userRequest.Password == "" {
		log.Println(utils.ERROR_REQUIRED)
		return http.StatusBadRequest, errors.New(utils.ERROR_REQUIRED)
	}

	// Fetch the user from the database
	statusCode, user, err := database.ReadUser(userRequest.Username)
	if err != nil {
		log.Println(utils.ERROR_USER_NOT_FOUND)
		return statusCode, err
	}

	// Check if the password is correct
	if !utils.CheckPassword(userRequest.Password, user.Password) {
		log.Println(utils.ERROR_PASSWORD_INCORRECT)
		return http.StatusUnauthorized, errors.New(utils.ERROR_PASSWORD_INCORRECT)
	}

	log.Println("User '" + userRequest.Username + "' successfully authenticated.")
	return http.StatusOK, nil
}

// GetGeoCode fetches the geocode for a given location
func (db *MongoDB) GetGeoCode(country utils.Country, city string) (int, utils.Coordinates, error) {

	var response utils.GeoCodeResults

	// Encode the city to be URL safe ('æ', 'ø', 'å' fixes)
	city = url.QueryEscape(city)

	// Fetch the API response from the city
	geoGet, err := http.Get(utils.GEOCODING_API + city)
	if err != nil {
		log.Println(utils.ERROR_GEOCODE)
		return http.StatusInternalServerError, utils.Coordinates{}, errors.New(utils.ERROR_GEOCODE)
	}
	defer geoGet.Body.Close()

	// Decode the response into the response struct
	err = json.NewDecoder(geoGet.Body).Decode(&response)
	if err != nil {
		log.Println(utils.ERROR_DECODING)
		return http.StatusInternalServerError, utils.Coordinates{}, errors.New(utils.ERROR_DECODING)
	}

	// Attempt to find the coordinates of the specified city
	location, found := utils.GetCity(response, country.IsoCode)

	// If the city can't be found in the specified country, we find the coordinates of the country
	if !found {
		location, err = utils.GetCountry(country.Name)
		if err != nil {
			log.Println(utils.ERROR_FETCH)
			return http.StatusServiceUnavailable, utils.Coordinates{}, errors.New(utils.ERROR_FETCH)
		}
	}

	// Create a return struct with the coordinates
	coordinates := utils.Coordinates{
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
	}

	return http.StatusOK, coordinates, nil
}

// GetWeather fetches the weather data for a given location
func (db *MongoDB) GetWeather(country utils.Country, city string) (int, utils.WeatherData, error) {

	// Fetch the geocode for the location
	statusCode, coordinates, err := db.GetGeoCode(country, city)
	if err != nil {
		return statusCode, utils.WeatherData{}, err
	}

	// Generate the weather URL
	weatherURL := utils.GenerateWeatherURL(coordinates)

	// Fetch the weather data from the API
	weatherGet, err := http.Get(weatherURL)
	if err != nil {
		log.Println(utils.ERROR_FETCH)
		return http.StatusServiceUnavailable, utils.WeatherData{}, errors.New(utils.ERROR_FETCH)
	}

	// Decode the response into the response struct
	var response utils.WeatherData
	err = json.NewDecoder(weatherGet.Body).Decode(&response)
	if err != nil {
		log.Println(utils.ERROR_DECODING)
		return http.StatusInternalServerError, utils.WeatherData{}, errors.New(utils.ERROR_DECODING)
	}

	return http.StatusOK, response, nil
}

// SetWeeklyWeather sets the weekly weather
func (db *MongoDB) SetWeeklyWeather(weather utils.WeatherData) (utils.WeeklyWeather, error) {

	// Return an error if the hourly data is not complete
	if len(weather.Hourly.Time) != utils.WEEKLY_HOURS {
		return utils.WeeklyWeather{}, fmt.Errorf(utils.ERROR_FETCH)
	}

	// Initialize the daily and weekly weather structs
	weeklyWeather := utils.WeeklyWeather{}
	dailyWeather := utils.DailyWeather{}

	// Set the current weather
	weeklyWeather.Today = utils.SetCurrentWeather(weather)

	// Initialize variable to count the remaining hours
	hour := 24

	// Iterate over the next seven days
	for i := 0; i < 6; i++ {
		dailyWeather.Date = utils.ExtractDate(weather.Hourly.Time[hour])

		// Iterate over every hour in each day
		for j := 0; j < 24; j++ {
			hourlyWeather := utils.HourlyWeather{
				Hour:          utils.ExtractHour(weather.Hourly.Time[hour]),
				Temperature:   weather.Hourly.Temperature[hour],
				Precipitation: weather.Hourly.Precipitation[hour],
				CloudCover:    weather.Hourly.CloudCover[hour],
				WindSpeed:     weather.Hourly.WindSpeed[hour],
			}
			hourlyWeather.Condition = utils.DetermineWeatherCondition(hourlyWeather)
			dailyWeather.Hours = append(dailyWeather.Hours, hourlyWeather)
			hour++
		}
		weeklyWeather.RestOfWeek = append(weeklyWeather.RestOfWeek, dailyWeather)
	}

	return weeklyWeather, nil
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
		log.Println(utils.ERROR_FIELDS_REQUIRED)
		return http.StatusBadRequest, errors.New(utils.ERROR_FIELDS_REQUIRED)
	}

	// Set username and email to lowercase
	user.Username = strings.ToLower(user.Username)
	user.Email = strings.ToLower(user.Email)

	// Check if the username or email already exists
	if _, exists := m.users[user.Username]; exists {
		log.Println(utils.ERROR_EXISTS)
		return http.StatusBadRequest, errors.New(utils.ERROR_EXISTS)
	}

	// Enforce a password only containing characters '1234567890'
	if !utils.EnforcePassword(user.Password) {
		log.Println(utils.LOG_PASSWORD_INVALID)
		return http.StatusBadRequest, errors.New(utils.ERROR_PASSWORD_INVALID)
	}

	var err error

	// Hash the password
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		log.Println(utils.ERROR_PASSWORD_HASH)
		return http.StatusInternalServerError, errors.New(utils.ERROR_PASSWORD_HASH)
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

	// Check if user exists, if not, then we assume a change of username was attempted
	currentValue, exists := m.users[username]
	if !exists {
		return http.StatusBadRequest, fmt.Errorf("user %s does not exist/username can't be changed", user.Username)
	}

	// Disallow any attempted change of username
	if user.Username != currentValue.Username {
		return http.StatusBadRequest, errors.New(utils.ERROR_USERNAME_CHANGE)
	}

	// Check if the password is changed
	if user.Password != "" || user.Password != currentValue.Password {

		// Apply constraints and hash the password if it is changed
		if !utils.EnforcePassword(user.Password) {
			return http.StatusBadRequest, errors.New(utils.ERROR_PASSWORD_INVALID)
		}

		if len(user.Password) < 8 {
			return http.StatusBadRequest, errors.New(utils.ERROR_PASSWORD_LENGTH)
		}

		var err error
		// Hash the password
		user.Password, err = utils.HashPassword(user.Password)
		if err != nil {
			return http.StatusInternalServerError, errors.New(utils.ERROR_PASSWORD_HASH)
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

// AuthenticateUser authenticates a user in the database
func (m *MockDB) AuthenticateUser(database Database, userRequest utils.UserAuthentication) (int, error) {

	// Check required fields
	if userRequest.Username == "" || userRequest.Password == "" {
		log.Println(utils.ERROR_REQUIRED)
		return http.StatusBadRequest, errors.New(utils.ERROR_REQUIRED)
	}

	// Fetch the user from the database
	statusCode, user, err := database.ReadUser(userRequest.Username)
	if err != nil {
		return statusCode, err
	}

	// Check if the password is correct
	if !utils.CheckPassword(userRequest.Password, user.Password) {
		log.Println(utils.ERROR_PASSWORD_INCORRECT)
		return http.StatusUnauthorized, errors.New(utils.ERROR_PASSWORD_INCORRECT)
	}

	log.Println("User '" + userRequest.Username + "' successfully authenticated.")
	return http.StatusOK, nil
}

// GetGeoCode fetches the geocode for a given location
func (m *MockDB) GetGeoCode(country utils.Country, city string) (int, utils.Coordinates, error) {

	jsonFile := utils.ParseFile(utils.GEOLOCATIONS_TEST_FILE)

	// Unmarshal the file to a response struct
	response := utils.GeoCodeResults{}
	err := json.Unmarshal(jsonFile, &response)
	if err != nil {
		log.Println(utils.ERROR_DECODING)
		return http.StatusInternalServerError, utils.Coordinates{}, errors.New(utils.ERROR_DECODING)
	}

	// Attempt to find the coordinates of the specified city (GetCountry is not included in the test)
	location, found := utils.GetCity(response, country.IsoCode)

	// Capture edge case where city is not found in the given country
	if !found {
		log.Println(utils.ERROR_COUNTRY_NOT_FOUND)
		return http.StatusNotFound, utils.Coordinates{}, errors.New(utils.ERROR_COUNTRY_NOT_FOUND)
	}

	// Create a return struct with the coordinates
	coordinates := utils.Coordinates{
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
	}

	return http.StatusOK, coordinates, nil
}

// GetWeather fetches the weather data for a given location
func (m *MockDB) GetWeather(country utils.Country, city string) (int, utils.WeatherData, error) {

	jsonFile := utils.ParseFile(utils.WEATHER_TEST_FILE)

	// Unmarshal the file to a response struct
	response := utils.WeatherData{}
	err := json.Unmarshal(jsonFile, &response)
	if err != nil {
		log.Println(utils.ERROR_DECODING)
		return http.StatusInternalServerError, utils.WeatherData{}, errors.New(utils.ERROR_DECODING)
	}

	return http.StatusOK, response, nil
}

// SetWeeklyWeather sets the weekly weather
func (m *MockDB) SetWeeklyWeather(weather utils.WeatherData) (utils.WeeklyWeather, error) {

	// Check if the number of hours in the weather data is correct
	if len(weather.Hourly.Time) != utils.WEEKLY_HOURS {
		return utils.WeeklyWeather{}, errors.New(utils.ERROR_DECODING)
	}

	// Create the weekly weather struct
	weeklyWeather := utils.WeeklyWeather{}
	weeklyWeather.Today.Date = utils.ExtractDate(weather.Hourly.Time[0])

	// Variable to count the hours through the week
	hour := 0

	// Set the date for the first day
	weeklyWeather.Today.Date = utils.ExtractDate(weather.Hourly.Time[hour])

	// Set values for the first day
	for i := 0; i < 24; i++ {
		hourlyWeather := utils.HourlyWeather{
			Hour:          weather.Hourly.Time[hour],
			Temperature:   weather.Hourly.Temperature[hour],
			Precipitation: weather.Hourly.Precipitation[hour],
			CloudCover:    weather.Hourly.CloudCover[hour],
			WindSpeed:     weather.Hourly.WindSpeed[hour],
		}
		hourlyWeather.Condition = utils.DetermineWeatherCondition(hourlyWeather)
		weeklyWeather.Today.Hours = append(weeklyWeather.Today.Hours, hourlyWeather)
		hour++
	}

	// Create the daily weather struct
	dailyWeather := utils.DailyWeather{}

	// Set values for the rest of the week
	for i := 0; i < 6; i++ {

		// Set the date for each day
		dailyWeather.Date = utils.ExtractDate(weather.Hourly.Time[hour])

		// Iterate over every hour in each day
		for j := 0; j < 24; j++ {
			hourlyWeather := utils.HourlyWeather{
				Hour:          weather.Hourly.Time[hour],
				Temperature:   weather.Hourly.Temperature[hour],
				Precipitation: weather.Hourly.Precipitation[hour],
				CloudCover:    weather.Hourly.CloudCover[hour],
				WindSpeed:     weather.Hourly.WindSpeed[hour],
			}
			hourlyWeather.Condition = utils.DetermineWeatherCondition(hourlyWeather)
			dailyWeather.Hours = append(dailyWeather.Hours, hourlyWeather)
			hour++
		}
		weeklyWeather.RestOfWeek = append(weeklyWeather.RestOfWeek, dailyWeather)
	}

	return weeklyWeather, nil
}
