package utils

// UserRegistration is a struct for user registration
type UserRegistration struct {
	Username   string          `json:"username"` // Enforced unique in the database
	Password   string          `json:"password,omitempty"`
	Email      string          `json:"email,omitempty"` // Enforced unique in the database
	Country    Country         `json:"country"`
	City       string          `json:"city"`
	Preference UserPreferences `json:"preferences,omitempty"`
}

// Country is a struct for holding country information
type Country struct {
	Name    string `json:"label"`
	IsoCode string `json:"value"`
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

// Coordinates is a struct for holding latitude and longitude
type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GeoCodeResults struct {
	Results []GeoCodeResponse `json:"results"`
}

// GeoCodeResponse is a struct for the response in the geocode API
type GeoCodeResponse struct {
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	CountryCode string  `json:"country_code"`
}

// WeatherResponse is a struct for the response in the weather endpoint
type WeatherResponse struct {
}
