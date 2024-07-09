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

// REGISTRATIONS is the endpoint for user registration
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

// SLASH is the slash character
export const SLASH = "/"

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


// LOCAL STORAGE
// LOGGEDIN is the key for the logged in status in local storage
export const LOGGEDIN = "loggedIn"

export const USERNAME = "username"

// VALUE_TRUE is the value for true in local storage
export const VALUE_TRUE = "true"

// VALUE_FALSE is the value for false in local storage
export const VALUE_FALSE = "false"


// TODO Move constants over to this file