package utils

// UserRegistration is a struct for user registration
type UserRegistration struct {
	Username   string          `json:"username"`
	Password   string          `json:"password"`
	Email      string          `json:"email"`
	Preference UserPreferences `json:"preferences"`
}

// UserPreferences is a struct for user preferences
type UserPreferences struct {
	Football bool `json:"football",omitempty`
	Movies   bool `json:"movies",omitempty`
	Weather  bool `json:"weather",omitempty`
}
