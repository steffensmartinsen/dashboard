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

// AUTHENTICATION_PATH is the path to the authentication endpoint
const AUTHENTICATION_PATH = "auth"

// AUTHENTICATION_ENDPOINT is the endpoint for authentication
const AUTHENTICATION_ENDPOINT = DEFAULT_PATH + AUTHENTICATION_PATH + "/"

// COLLECTION_USERS is the MongoDB collection for users
const COLLECTION_USERS = "users"
