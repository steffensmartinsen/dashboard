# Goal of the project
The goal of the project is to show of some backend and frontend skills, and combining them to create a coherent service. The service utilizes ReactJS frontend, which I learned mostly through this project, a Go backend and a MongoDB document database. 

# WORK IN PROGRESS

## Planned implementations
* Weather report for a given location. Fetched and calculated in the backend, and displayed in the frontend.
* Football statistics for a given team.
* Delete user

For the time being, this project is not prioritized. 

## Current functionality:
Users can register accounts with a specified country and city, and choose preferences in terms of movie recommendation, favourite team and weather. When the user enters the home page it displays the current weather in the given city they picked (or country if the service cannot find the city). The user can edit their password and change account settings. 

--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
# Table of Contents

- [Dashboard](#dashboard)
- [Endpoints](#endpoints)
  - [Registrations](#registrations)
    - [POST](#post)
    - [GET](#get)
    - [PUT](#put)
    - [DELETE](#delete)
  - [Authentication](#authentication)
    - [POST](#post-1)
  - [Weather](#weather)
    - [GET](#get-1)
- [Cookie Handling](#cookie-handling)
  - [set-cookie](#set-cookie)
  - [get-cookie](#get-cookie)
  - [delete-cookie](#delete-cookie)
- [Test Coverage](#test-coverage)

# Dashboard
Welcome to this Dashboard API. The service allows you to register an account, and specify which features you would like to see on your dashboard. The service will then provide you with the requested information.

From the [PROG2005](https://www.ntnu.edu/studies/courses/PROG2005#tab=omEmnet) course I already have experience with Firestore and XAMPP SQL database impelemtations. Therefore, I wanted to test MongoDB in this project.

A simple frontend serving the backend API is built using React. I am using this project as an opportunity to learn React, and the frontend is developed alongside the backend functionality. The frontend is not a priority, and the main focus is on the backend API.

# Endpoints

## Registrations
Endpoint url: `http://localhost:8080/dashboard/v1/registrations/`

The *Registrations* endpoint allows for account creation, deletion, changes, and for the service to retrieve meta-information about a user in the system.<br>
A user's username and email is enforced unique by the MongoDB database. The password can only contain numbers.<br>
An account registers with it's username, email, password, and preferences as to what to display on the dashboard.

Supported methods: `POST`, `GET`, `PUT`, `DELETE`

### POST
The `POST` request to the registrations endpoint allows users to register an account. <br>
The request body must contain the `username`, `password` and `email` fields. <br>
The `password` field must contain all numbers to ensure no proper password is used in this dummy service.
The password is hashed before it is stored in the database.
<br>
The request body also contains boolean values for what the user wishes to have available on their dashboards. The `preferences` field is optional, and if omitted, the user will have no preferences set.

**Invocation URL:** `http://localhost:8080/dashboard/v1/registrations/`<br>
**Method:** `POST` <br>
**Success Response:** `201 Created`

**Example JSON request body:**
```
{
    "username": "user",
    "password": "123456",
    "email": "user@fakemail.com"
    "preferences": {
        "weather": true,
        "movies": true,
        "football": true
    }
}
```


### GET
The `GET` request to the registrations endpoint allows users to retrieve their account information.<br>
The **invocation URL** must specify the user's username, which is enforced unique by the MongoDB database.

**Invocation URL:** `http://localhost:8080/dashboard/v1/registrations/{username}`<br>
**Method:** `GET` <br>
**Success Response:** `200 OK`

**Example Invocation URL:** `http://localhost:8080/dashboard/v1/registrations/user`<br>
**Example Response Body:**
```
{
{
    "username": "user",
    "password": "0987654321",
    "email": "user@fakemail.com"
    "preferences": {
        "weather": true,
        "movies": true,
        "football": false
    }
}
```

### PUT
The `PUT` request allows users to change their account information.<br>
The `PUT` request body must contain the username, which is enforced unique by the MongoDB database.
The endpoint will read the username from the request body.<br>
Users can change every value in their account, except for the username.<br>
Additionally, if the fields in the `preferences` object are omitted, the user's preferences will be set to `false`.

**Invocation URL:** `http://localhost:8080/dashboard/v1/registrations/`<br>
**Method:** `PUT` <br>
**Success Response:** `200 OK`

**Example Invocation URL:** `http://localhost:8080/dashboard/v1/registrations/`<br>
**Example JSON request body:**
```
{
    "username": "user",
    "password": "0987654321",
    "email": "
    "preferences": {
        "weather": true,
        "movies": false,
        "football": true
    }
}
```

### DELETE
The `DELETE` request to the registrations endpoint allows users to delete their account.<br>
The **invocation URL** must specify the user's username, which is enforced unique by the MongoDB database.<br>
A successful deletion of a user account returns a `204 No Content` and an empty body.

**Invocation URL:** `http://localhost:8080/dashboard/v1/registrations/{username}`<br>
**Method:** `DELETE` <br>
**Success Response:** `204 No Content`

**Example Invocation URL:** `http://localhost:8080/dashboard/v1/registrations/user`<br>

## Authentication
Endpoint url: `http://localhost:8080/dashboard/v1/auth/`

The **Authentication** endpoint authenticates a user's login credentials.<br>
The endpoint is used to verify a user's login credentials, and if successful, the user is granted access to the dashboard.

Supported methods: `POST`

### POST
The `POST` request to the authentication endpoint checks whether the user's login credentials match the ones stored in the database.<br>
The request body must contain the `username` and `password` fields.

**Invocation URL:** `http://localhost:8080/dashboard/v1/auth/`<br>
**Method:** `POST` <br>
**Success Response:** `200 OK`

**Example JSON request body:**
```
{
    "username": "user",
    "password": "1234567890"
}
```

## Weather
Endpoint URL: `http://localhost:8080/dashboard/v1/weather/{username}`

The **Weather** endpoint fetches the weather forecast for a specified user supplied from the user's registered city and country. 
The endpoint fetches data from the *Weather Forecast API* based on the user's location.
To find the location for the weather forecast, the endpoint first fetches the coordinates of the city/country from the *Geocoding API*.<br>
The endpoint separates the current day from the rest of the week, where today returns the current hour and the remaining hours of the day, 
whereas the next 6 days are separated by date, and returns 24 hours each.<br>
The endpoint utilizes an internal function that maps the date and hour to the corresponding weather data for that given hour. 

The internal function also determines a string suited for the given weather (e.g. *cloudy*, *clear*, *partly cloudy* etc.), 
determined on the basis of the data for a given hour.<br>
The response body contains every hour for the next week with the given string condition, in addition to temperature, 
precipitation, cloud cover, and wind speed.<br>
The endpoint supports `GET` requests and enforces the user's username to be specified in the URL.<br>

### GET
The `GET` request to the weather endpoint fetches the weather forecast for the user's registered city and country.<br>

**Invocation URL:** `http://localhost:8080/dashboard/v1/weather/{username}`<br>
**Method:** `GET` <br>
**Success Response:** `200 OK`

**Example Invocation URL:** `http://localhost:8080/dashboard/v1/weather/user`<br>
**Example Response Body:** (29-06-2024 23:00)
```
{
    "today": {
        "date": "2024-06-29",
        "hours": [
            {
                "hour": "23",
                "condition": "mostly clear",
                "temperature": 14,
                "precipitation": 0,
                "cloudCover": 15,
                "windSpeed": 3
            }
        ]
    },
    "restOfWeek": [
        {
            "date": "2024-06-29",
            "hours": [
                {
                    "hour": "00",
                    "condition": "rainy",
                    "temperature": 15.2,
                    "precipitation": 0.1,
                    "cloudCover": 92,
                    "windSpeed": 5.4
                },
                {
                    "hour": "01",
                    "condition": "mostly sunny",
                    "temperature": 14.6,
                    "precipitation": 0,
                    "cloudCover": 34,
                    "windSpeed": 4.4
                },
                
            ...
```


# Cookie Handling
The service uses cookies to store the user's session information. The cookies are set when a user logs in, and are deleted when the user logs out.

## set-cookie
**Invocation URL:** `http://localhost:8080/dashboard/v1/set-cookie/` <br>
**METHOD:** `POST`

A `POST` request is sent to the endpoint containing the username of the user requesting a cookie. The API generates a random token to set as the cookies value. 
The cookie is set with a 24-hour expiration time and the `HttpOnly` flag set to `true`.
Successful setting of the cookie returns a `200 OK` status code.

## get-cookie
**Invocation URL:** `http://localhost:8080/dashboard/v1/get-cookie/{username}` <br>
**METHOD:** `GET`

A `GET` request is sent to the endpoint containing the username of the user requesting the cookie. The API retrieves the cookie value from server memory and returns the generated token in the response body.
If the user does not have a cookie set, the API returns a `404 Not Found` status code.
Successful retrieval of the cookie returns a `200 OK` status code.

## delete-cookie
**Invocation URL:** `http://localhost:8080/dashboard/v1/delete-cookie/{username}` <br>
**METHOD:** `DELETE`

A `DELETE` request is sent to the endpoint containing the username of the user requesting the cookie deletion. The API deletes the cookie from server memory.
This endpoint is automatically called with the username when the user logs out of the frontend. If a user does not have a cookie set, the API returns a `404 Not Found` status code.
Successful deletion of the cookie returns a `204 No Content` status code.

# Dependencies
The service depends on the following endpoints to fetch relevant information:
- [Geocoding API](https://open-meteo.com/en/docs/geocoding-api)
- [Weather Forecast API](https://open-meteo.com/en/docs)

# Test Coverage
All HTTP methods on the registration endpoint are covered with tests inside `endpoints/registrations_test.go`. 
The `POST` method on the authorisation endpoint is covered with tests inside `endpoints/authorisationHandler_test.go`.<br>
The `GET` request on the weather endpoint is covered with test stubbing the *Weather Forecast API* inside `endpoints/weatherHandler_test.go`.<br>
To run the test on the **registrationHandler.go** and **authorisationHandler.go** files, run `go test` inside the `endpoints` folder from the terminal.<br>


All the CRUD functions supplementing the registration endpoint, alongside a stubbed API test of the GetGeoCode function, are covered with tests inside `database/databaseOperations_test.go`<br>
To run the test on the **databaseOperations.go** file, run `go test` inside the `database` folder from the terminal.<br>