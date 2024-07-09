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

const AUTH = "/auth/"

// SLASH is the slash character
export const SLASH = "/"

// ENDPOINT_REGISTRATIONS is the URL for user registration
export const ENDPOINT_REGISTRATIONS = LOCALHOST + PORT + DASHBOARDS + VERSION + REGISTRATIONS

// ENDPOINT_AUTHENTICATION is the URL for user authentication
export const ENDPOINT_AUTH = LOCALHOST + PORT + DASHBOARDS + VERSION + AUTH


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

// TODO Move constants over to this file