// Weather Report
// CONDITION_RAINY is the condition for rainy weather used to determine weather icon
export const CONDITION_RAINY = "Rainy"

// CONDITION_CLOUDY is the condition for cloudy weather used to determine weather icon
export const CONDITION_CLOUDY = "Cloudy"

// CONDITION_MOSTLY_CLOUDY is the condition for mostly cloudy weather used to determine weather icon
export const CONDITION_MOSTLY_CLOUDY = "Mostly Cloudy"

// CONDITION_PARTLY_CLOUDY is the condition for partly cloudy weather used to determine weather icon
export const CONDITION_PARTLY_CLOUDY = "Partly Cloudy"

// CONDITION_MOSTLY_SUNNY is the condition for mostly sunny weather used to determine weather icon
export const CONDITION_MOSTLY_SUNNY = "Mostly Sunny"

// CONDITION_MOSTLY_CLEAR is the condition for mostly clear weather used to determine weather icon
export const CONDITION_MOSTLY_CLEAR = "Mostly Clear"

// CONDITION_CLEAR_DAY is the condition for clear day weather used to determine weather icon
export const CONDITION_CLEAR_DAY = "Clear Day"


// URLS
// LOCALHOST is the base URL for the backend server
const LOCALHOST = "http://localhost:"

// PORT is the port number for the backend server
const PORT = "8080"

// DASHBOARDS is the base of the backend server URL for the service
const DASHBOARDS = "/dashboards/"

// VERSION is the version of the backend server
export const VERSION = "v1"

// REGISTRATIONS is the endpoint for user registration in the backend
const REGISTRATIONS = "/registrations/"

// AUTH is the endpoint for user authentication
const AUTH = "/auth/"

// SET_COOKIE is the endpoint for setting a cookie
const SET_COOKIE = "/set-cookie/"

// GET_COOKIE is the endpoint for getting a cookie
const GET_COOKIE = "/get-cookie/"

// DELETE_COOKIE is the endpoint for deleting a cookie
const DELETE_COOKIE = "/delete-cookie/"

// WEATHER is the endpoint for getting weather data
export const WEATHER = "/weather/"

// LOGIN is the endpoint for user login
export const LOGIN = "/login"

// REGISTER is the endpoint for user registration on the frontend
export const REGISTER = "/register"

// SLASH is the slash character
export const SLASH = "/"

// ROOT is the root URL
export const ROOT = "/"

// ENDPOINT_REGISTRATIONS is the URL for user registration
export const ENDPOINT_REGISTRATIONS = LOCALHOST + PORT + DASHBOARDS + VERSION + REGISTRATIONS

// ENDPOINT_AUTHENTICATION is the URL for user authentication
export const ENDPOINT_AUTH = LOCALHOST + PORT + DASHBOARDS + VERSION + AUTH

// ENDPOINT_SET_COOKIE is the URL for setting a cookie
export const ENDPOINT_SET_COOKIE = LOCALHOST + PORT + DASHBOARDS + VERSION + SET_COOKIE

// ENDPOINT_GET_COOKIE is the URL for getting a cookie
export const ENDPOINT_GET_COOKIE = LOCALHOST + PORT + DASHBOARDS + VERSION + GET_COOKIE

// ENDPOINT_DELETE_COOKIE is the URL for deleting a cookie
export const ENDPOINT_DELETE_COOKIE = LOCALHOST + PORT + DASHBOARDS + VERSION + DELETE_COOKIE

// ENDPOINT_WEATHER is the URL for getting weather data
export const ENDPOINT_WEATHER = LOCALHOST + PORT + DASHBOARDS + VERSION + WEATHER

//HTTP Methods and Headers
// GET is the HTTP GET method
export const METHOD_GET = "GET"

// POST is the HTTP POST method
export const METHOD_POST = "POST"

// PUT is the HTTP PUT method
export const METHOD_PUT = "PUT"

// DELETE is the HTTP DELETE method
export const METHOD_DELETE = "DELETE"

// HEADER_APPLICATION_JSON is the application/json header
export const HEADER_APPLICATION_JSON = "application/json"

// INCLUDE_CREDENTIALS is the include credentials header
export const INCLUDE_CREDENTIALS = "include"

// EXTERNAL_GITHUB is the URL for the external Github repository
export const EXTERNAL_GITHUB = "https://github.com/steffensmartinsen/dashboard"


// LOCAL STORAGE
// LOGGEDIN is the key for the logged in status in local storage
export const LOGGEDIN = "loggedIn"

export const USERNAME = "username"

// VALUE_TRUE is the value for true in local storage
export const VALUE_TRUE = "true"

// VALUE_FALSE is the value for false in local storage
export const VALUE_FALSE = "false"


// AUTHENTICATION
// ACCEPTED_PW_VALUES is the accepted password values
export const ACCEPTED = ["1", "2", "3", "4", "5", "6", "7", "8", "9", "0"];


// ERROR MESSAGES

// ERROR_500 is the error message for a 500 Internal Server Error
export const ERROR_500 = "Internal Server Error"

// ERROR_USERNAME is the error message for an invalid username
export const ERROR_USERNAME = "Invalid username"

// ERROR_USERNAME_REQUIRED is the error message for a required username
export const ERROR_USERNAME_REQUIRED = "Username is required"

// ERROR_USERNAME_LOWER is the error message for a username that is not lowercase
export const ERROR_USERNAME_LOWER = "Username must be lowercase"

// ERROR_EXISTS is the error message for a username or email that already exists
export const ERROR_EXISTS = "Username or e-mail already exists"

// ERROR_EMAI_EXISTS is the error message for an email that already exists
export const ERROR_EMAIL_EXISTS = "E-mail already exists"

// ERROR_EMAIL_REQUIRED is the error message for a required email
export const ERROR_EMAIL_REQUIRED = "Email is required"

// ERROR_EMAIL_INVALID
export const ERROR_EMAIL_INVALID = "Invalid email"

// ERROR_COOKIE_SET
export const ERROR_COOKIE_SET = "Error setting cookie"

// ERROR_COOKIE_GET
export const ERROR_COOKIE_GET = "Error getting cookie"

// ERROR_COOKIE_DELETE
export const ERROR_COOKIE_DELETE = "Error deleting cookie"

// ERROR_SELECT_COUNTRY is the error message for a required country
export const ERROR_SELECT_COUNTRY = "Please select a country"

// ERROR_SELECT_CITY is the error message for a required city
export const ERROR_SELECT_CITY = "Please enter a city"

// ERROR_PASSWORD_REQUIRED is the error message for a required password
export const ERROR_PASSWORD_REQUIRED = "Password is required"

// ERROR_PASSWORD_LENGTH is the error message for a password that is too short
export const ERROR_PASSWORD_LENGTH = "Password must be at least 8 characters"

// ERROR_PASSWORD_MATCH is the error message for a password that does not match
export const ERROR_PASSWORD_MATCH = "Passwords do not match"

// ERROR_PASSWORD_INVALID is the error message for a password containing illegal characters
export const ERROR_PASSWORD_INVALID = "Wrong password. Please don't use an actual password for this. The only accepted characters are '1234567890'"

// ERROR_PASSWORD_INCORRECT is the error message for an incorrect password
export const ERROR_PASSWORD_INCORRECT = "Incorrect password"

// ERROR_WEATHER is the error message for a weather error
export const ERROR_WEATHER = "Error getting weather data"

// ERROR_SELECT_TEAM is the error message for a team error
export const ERROR_SELECT_TEAM = "Please enter a team"

// ERROR_UNDEFINED is the error message for an undefined error
export const ERROR_UNDEFINED = "Something went wrong"

// ERROR_LOGIN is the error message for a login error
export const ERROR_LOGIN = "Invalid username or password"


// SUCCESS MESSAGES
// SUCCESS_USER_CREATED is the success message for a user created successfully
export const SUCCESS_USER_CREATED = "User created successfully"

// SUCCESS_USER_AUTHENTICATED is the success message for a user authenticated successfully
export const SUCCESS_USER_AUTHENTICATED = "User authenticated successfully"

// SUCCESS_USER_UPDATED is the success message for a user updated successfully
export const SUCCESS_USER_UPDATED = "User updated successfully"

// SUCCESS_COOKIE_SET is the success message for setting a cookie
export const SUCCESS_COOKIE_SET = "Cookie set successfully"

// SUCCESS_COOKIE_GET is the success message for getting a cookie
export const SUCCESS_COOKIE_GET = "Cookie retrieved successfully"

// SUCCESS_COOKIE_DELETE is the success message for deleting a cookie
export const SUCCESS_COOKIE_DELETE = "Cookie deleted successfully"

// SUCCESS_PASSWORD_UPDATED is the success message for updating a password
export const SUCCESS_PASSWORD_UPDATED = "Password updated successfully"

// TODO Move constants over to this file