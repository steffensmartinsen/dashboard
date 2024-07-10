package utils

// URLS and DB
// DEFAULT_PORT is the default port for the service
const DEFAULT_PORT = "8080"

// DEFAULT_PATH is the root path.
const DEFAULT_PATH = "/dashboards/" + VERSION + SLASH

// VERSION is the current version of the service.
const VERSION = "v1"

// LOCALHOST is the localhost address
const LOCALHOST = "localhost"

// ROOT is the root path
const SLASH = "/"

// ENDPOINT_REGISTRATIONS holds the name for the registration endpoint
const ENDPOINT_REGISTRATIONS = "registrations"

// PATH_REGISTRATIONS is the path to the registration endpoint
const PATH_REGISTRATIONS = DEFAULT_PATH + ENDPOINT_REGISTRATIONS + SLASH

// PATH_AUTHENTICATION is the path to the authentication endpoint
const ENDPOINT_AUTHENTICATION = "auth"

// ENDPOINT_AUTHENTICATION is the endpoint for authentication
const PATH_AUTHENTICATION = DEFAULT_PATH + ENDPOINT_AUTHENTICATION + SLASH

// PATH_DASHBOARD is the path to the dashboard
const ENDPOINT_WEATHER = "weather"

// ENDPOINT_DASHBOARD is the endpoint for the dashboard
const PATH_WEATHER = DEFAULT_PATH + ENDPOINT_WEATHER + SLASH

// COLLECTION_USERS is the MongoDB collection for users
const COLLECTION_USERS = "users"

// ENDPOINT_SET_COOKIE is the endpoint for setting a cookie
const ENDPOINT_SET_COOKIE = "set-cookie"

// PATH_SET_COOKIE is the path to the set-cookie endpoint
const PATH_SET_COOKIE = DEFAULT_PATH + ENDPOINT_SET_COOKIE + SLASH

// ENDPOINT_GET_COOKIE is the endpoint for getting a cookie
const ENDPOINT_GET_COOKIE = "get-cookie"

// PATH_GET_COOKIE is the path to the get-cookie endpoint
const PATH_GET_COOKIE = DEFAULT_PATH + ENDPOINT_GET_COOKIE + SLASH

// ENDPOINT_DELETE_COOKIE is the endpoint for deleting a cookie
const ENDPOINT_DELETE_COOKIE = "delete-cookie"

// PATH_DELETE_COOKIE is the path to the delete-cookie endpoint
const PATH_DELETE_COOKIE = DEFAULT_PATH + ENDPOINT_DELETE_COOKIE + SLASH

// GEOCODING_API is the URL for the geo-coding API
const GEOCODING_API = "https://geocoding-api.open-meteo.com/v1/search/?name="

// GEOLOCATIONS_TEST_FILE is the local path to the test file for geolocations
const GEOLOCATIONS_TEST_FILE = "../test_data/geolocations.json"

// WEATHER_TEST_FILE is the local path to the test file for weather
const WEATHER_TEST_FILE = "../test_data/weather.json"

// WEATHER_API_BASE is the base URL for the weather API
const WEATHER_API_BASE = "https://api.open-meteo.com/v1/forecast?"

// WEATHER_API_LAT is the latitude parameter for the weather API
const WEATHER_API_LAT = "latitude="

// WEATHER_API_LON is the longitude parameter for the weather API
const WEATHER_API_LON = "&longitude="

// WEATHER_API_HOURLY is the hourly parameters for the weather API
const WEATHER_API_PARAMETERS = "&hourly=temperature_2m,precipitation,cloud_cover,wind_speed_10m&wind_speed_unit=ms"


// WEATHER HANDLING
// CONDITION_RAINY is the condition for rainy weather returned by the internal weather determiner
const CONDITION_RAINY = "Rainy"

// CONDITION_CLOUDY is the condition for cloudy weather returned by the internal weather determiner
const CONDITION_CLOUDY = "Cloudy"

// CONDITION_MOSTLY_CLOUDY is the condition for mostly cloudy weather returned by the internal weather determiner
const CONDITION_MOSTLY_CLOUDY = "Mostly Cloudy"

// CONDITION_PARTLY_CLOUDY is the condition for partly cloudy weather returned by the internal weather determiner
const CONDITION_PARTLY_CLOUDY = "Partly Cloudy"

// CONDITION_MOSTLY_SUNNY is the condition for mostly sunny weather returned by the internal weather determiner
const CONDITION_MOSTLY_SUNNY = "Mostly Sunny"

// CONDITION_MOSTLY_CLEAR is the condition for mostly clear weather returned by the internal weather determiner
const CONDITION_MOSTLY_CLEAR = "Mostly Clear"

// CONDITION_CLEAR_DAY is the condition for clear day weather returned by the internal weather determiner
const CONDITION_CLEAR_DAY = "Clear Day"

// TIME_FORMAT is the format for the time string returned in the weather handler
const TIME_FORMAT = "2006-01-02T15:04"

// WEEKLY_HOURS is the total number of hours in a week
const WEEKLY_HOURS = 168

// USERNAME is the constant value for the username field
const USERNAME = "username"

// BSON_SET is the constant value for the $set operator in MongoDB
const BSON_SET = "$set"


// ERROR MESSAGES
// ERROR_FIELDS_REQUIRED is the error message for missing fields in the registration request
const ERROR_FIELDS_REQUIRED = "username, password, and email are required fields"

// ERROR_EXISTS is the error message for a user already existing in the database
const ERROR_EXISTS = "username or email already exists"

// ERROR_PASSWORD_INVALID is the error message for a password containing illegal characters
const ERROR_PASSWORD_INVALID = "please don't use an actual password for this. The only accepted characters are '1234567890'"

// LOG_PASSWORD_INVALID is the log message for a password containing illegal characters
const LOG_PASSWORD_INVALID = "illegal password characters"

// ERROR_PASSWORD_HASH is the error message for a password hash failure
const ERROR_PASSWORD_HASH = "error hashing password"

// ERROR_USER_NOT_FOUND is the error message for a user not found in the database
const ERROR_USER_NOT_FOUND = "user not found"

// ERROR_USERNAME_CHANGE is the error message for a username change attempt
const ERROR_USERNAME_CHANGE = "username cannot be changed"

// ERROR_USER_FETCH is the error message for a user fetch failure
const ERROR_USER_FETCH = "error fetching user"

// ERROR_EMAIL_EXISTS is the error message for an email already existing in the database
const ERROR_EMAIL_EXISTS = "email already exists"

// ERROR_PASSWORD_LENGTH is the error message for a password being too short
const ERROR_PASSWORD_LENGTH = "password must be at least 8 characters long"

// ERROR_USER_UPDATE is the error message for a user update failure
const ERROR_USER_UPDATE = "error updating user"

// ERROR_USER_DELETE is the error message for a user delete failure
const ERROR_USER_DELETE = "error deleting user"

// ERROR_NOT_PROVIDED is the error message for a missing field in the request
const ERROR_REQUIRED = "username and password required"

// ERROR_PASSWORD_INCORRECT is the error message for an incorrect password
const ERROR_PASSWORD_INCORRECT = "incorrect password"

// ERROR_GEOCODE is the error message for a geocode failure
const ERROR_GEOCODE = "error fetching geocode"

// ERROR_DECODING is the error message for a decoding failure
const ERROR_DECODING = "error decoding response"

// ERROR_FETCH is the error message for a country fetch failure
const ERROR_FETCH = "error fetching resource"

// ERROR_COUNTRY_NOT_FOUND is the error message for a country not found in the external API
const ERROR_COUNTRY_NOT_FOUND = "country not found"