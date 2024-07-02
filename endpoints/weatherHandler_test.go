package endpoints

import (
	"dashboard/database"
	"dashboard/utils"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	TEMPERATURE   = 15.7
	PRECIPITATION = 0.00
	CLOUD_COVER   = 100
	WIND_SPEED    = 7.8
)

func TestWeatherHandler(t *testing.T) {

	db := database.NewMockDB()
	client := http.Client{}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		WeatherHandler(db, w, r)
	}))
	defer server.Close()

	// Test case with user containing all required fields
	user := utils.UserRegistration{
		Username: USERNAME,
		Password: PASSWORD,
		Email:    EMAIL,
		Country:  utils.Country{Name: COUNTRY, IsoCode: ISOCODE},
		City:     CITY,
	}
	db.CreateUser(user)

	resp, err := client.Get(server.URL + utils.PATH_WEATHER + user.Username + "/")
	if err != nil {
		t.Fatal("Get request to URL failed:", err.Error())
	}

	// Check the status code
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("expected %d but got %d", http.StatusOK, resp.StatusCode)
	}

	// Check the response body
	var response utils.WeeklyWeather
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatal("Failed to decode response body:", err.Error())
	}

	// TODO Make tests runable

	if response.Today.Hours[0].Temperature != TEMPERATURE {
		t.Errorf("expected temperature %f but got %f", TEMPERATURE, response.Today.Hours[0].Temperature)
	}
	if response.Today.Hours[0].Precipitation != PRECIPITATION {
		t.Errorf("expected precipitation %f but got %f", PRECIPITATION, response.Today.Hours[0].Precipitation)
	}
	if response.Today.Hours[0].CloudCover != CLOUD_COVER {
		t.Errorf("expected cloud cover %d but got %f", CLOUD_COVER, response.Today.Hours[0].CloudCover)
	}
	if response.Today.Hours[0].WindSpeed != WIND_SPEED {
		t.Errorf("expected wind speed %f but got %f", WIND_SPEED, response.Today.Hours[0].WindSpeed)
	}
	if response.Today.Hours[8].Condition != utils.CONDITION_PARTLY_CLOUDY {
		t.Errorf("expected condition %s but got %s", utils.CONDITION_PARTLY_CLOUDY, response.Today.Hours[8].Condition)
	}

	log.Println("------- TestWeatherHandler passed -------")
}

//func setTestWeather(weather utils.WeatherData) (utils.WeeklyWeather, error) {
//
//	// Check if the number of hours in the weather data is correct
//	if len(weather.Hourly.Time) != utils.WEEKLY_HOURS {
//		return utils.WeeklyWeather{}, errors.New("invalid number of hours in the weather data")
//	}
//
//	// Create the weekly weather struct
//	weeklyWeather := utils.WeeklyWeather{}
//	weeklyWeather.Today.Date = utils.ExtractDate(weather.Hourly.Time[0])
//
//	// Variable to count the hours through the week
//	hour := 0
//
//	// Set the date for the first day
//	weeklyWeather.Today.Date = utils.ExtractDate(weather.Hourly.Time[hour])
//
//	// Set values for the first day
//	for i := 0; i < 24; i++ {
//		hourlyWeather := utils.HourlyWeather{
//			Hour:          weather.Hourly.Time[hour],
//			Temperature:   weather.Hourly.Temperature[hour],
//			Precipitation: weather.Hourly.Precipitation[hour],
//			CloudCover:    weather.Hourly.CloudCover[hour],
//			WindSpeed:     weather.Hourly.WindSpeed[hour],
//		}
//		hourlyWeather.Condition = utils.DetermineWeatherCondition(hourlyWeather)
//		weeklyWeather.Today.Hours = append(weeklyWeather.Today.Hours, hourlyWeather)
//		hour++
//	}
//
//	dailyWeather := utils.DailyWeather{}
//	// Set values for the rest of the week
//	for i := 0; i < 6; i++ {
//
//		// Set the date for each day
//		dailyWeather.Date = utils.ExtractDate(weather.Hourly.Time[hour])
//
//		// Iterate over every hour in each day
//		for j := 0; j < 24; j++ {
//			hourlyWeather := utils.HourlyWeather{
//				Hour:          weather.Hourly.Time[hour],
//				Temperature:   weather.Hourly.Temperature[hour],
//				Precipitation: weather.Hourly.Precipitation[hour],
//				CloudCover:    weather.Hourly.CloudCover[hour],
//				WindSpeed:     weather.Hourly.WindSpeed[hour],
//			}
//			hourlyWeather.Condition = utils.DetermineWeatherCondition(hourlyWeather)
//			dailyWeather.Hours = append(dailyWeather.Hours, hourlyWeather)
//			hour++
//		}
//		weeklyWeather.RestOfWeek = append(weeklyWeather.RestOfWeek, dailyWeather)
//	}
//
//	return weeklyWeather, nil
//}
