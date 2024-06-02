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
const ENDPOINT_REGISTRATIONS = DEFAULT_PATH + "registrations/"

// COLLECTION_USERS is the MongoDB collection for users
const COLLECTION_USERS = "users"
