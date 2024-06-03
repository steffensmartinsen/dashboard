# WORK IN PROGRESS

## Planned implementations
<ul>
  <li>Simple frontend home page, and other required pages for navigation. Low priority.</li>
  <li>Weather report for a given location.</li>
  <li>A recommendation system based on a given movie, though the selection of recommended movies will be limited by the number of movies included in the generated MongoDB sample collection.</li>
  <li>Football statistics for a given team.</li>
  <li>The service can register Webhooks for notifications for events happening with their football team</li>
</ul>

--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

# Dashboard
Welcome to this Dashboard API. The service allows you to register an account, and specify which features you would like to see on your dashboard. The service will then provide you with the requested information.

From the [PROG2005](https://www.ntnu.edu/studies/courses/PROG2005#tab=omEmnet) course I already have experience with Firestore and XAMPP SQL database impelemtations. Therefore, I wanted to test MongoDB in this project.

# Endpoints

## Registrations
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
**Method:** `POST`

**Example JSON request body:**
```
{
    "username": "user",
    "password": "123456",
    "email": "
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
**Method:** `GET`

**Example Invocation URL:** `http://localhost:8080/dashboard/v1/registrations/user`<br>
**Example Response Body:**
```
{
{
    "username": "user",
    "password": "0987654321",
    "email": "
    "preferences": {
        "weather": true,
        "movies": true,
        "football": false
    }
}
```

### PUT
The `PUT` request allows users to change their account information.<br>
The **invocation URL** must specify the user's username, which is enforced unique by the MongoDB database.<br>
Users can change every value in their account, except for the username.<br>
Additionally, if the fields in the `preferences` object are omitted, the user's preferences will be set to `false`.

**Invocation URL:** `http://localhost:8080/dashboard/v1/registrations/{username}`<br>
**Method:** `PUT`

**Example Invocation URL:** `http://localhost:8080/dashboard/v1/registrations/user`<br>
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
**Method:** `DELETE`

**Example Invocation URL:** `http://localhost:8080/dashboard/v1/registrations/user`<br>
