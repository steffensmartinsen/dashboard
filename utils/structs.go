package utils

// UserRegistration is a struct for user registration
type UserRegistration struct {
	Username   string          `json:"username"` // Enforced unique in the database
	Password   string          `json:"password,omitempty"`
	Email      string          `json:"email,omitempty"` // Enforced unique in the database
	Preference UserPreferences `json:"preferences,omitempty"`
}

// UserPreferences is a struct for user preferences
type UserPreferences struct {
	Football bool   `json:"football,omitempty"`
	Movies   bool   `json:"movies,omitempty"`
	Weather  bool   `json:"weather,omitempty"`
	Team     string `json:"team,omitempty"`
}

// UserAuthentication is a struct for user authentication
type UserAuthentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
