package utils

// MONGODB_URI is the MongoDB URI for the service
const MONGODB_URI = "mongodb+srv://viatheyboy:uvL3rwoDOrt42G8g@dashboards.ylq2uxl.mongodb.net/?retryWrites=true&w=majority&appName=dashboards"

// DEFAULT_PORT is the default port for the service
const DEFAULT_PORT = "8080"

// DEFAULT_PATH is the root path.
const DEFAULT_PATH = "/dashboards/" + VERSION + "/"

// VERSION is the current version of the service.
const VERSION = "v1"

// ENDPOINT_REGISTRATIONS holds the name for the registration endpoint
const ENDPOINT_REGISTRATIONS = "registrations"

// PATH_REGISTRATIONS is the path to the registration endpoint
const PATH_REGISTRATIONS = DEFAULT_PATH + ENDPOINT_REGISTRATIONS + "/"

// PATH_AUTHENTICATION is the path to the authentication endpoint
const ENDPOINT_AUTHENTICATION = "auth"

// ENDPOINT_AUTHENTICATION is the endpoint for authentication
const PATH_AUTHENTICATION = DEFAULT_PATH + ENDPOINT_AUTHENTICATION + "/"

// PATH_DASHBOARD is the path to the dashboard
const ENDPOINT_DASHBOARD = "dashboard"

// ENDPOINT_DASHBOARD is the endpoint for the dashboard
const PATH_DASHBOARD = DEFAULT_PATH + ENDPOINT_DASHBOARD + "/"

// COLLECTION_USERS is the MongoDB collection for users
const COLLECTION_USERS = "users"
