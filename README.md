# Dashboard
Welcome to this Dashboard API. The service allows you to register an account, and specify which features you would like to see on your dashboard. The service will then provide you with the requested information.

From the [PROG2005](https://www.ntnu.edu/studies/courses/PROG2005#tab=omEmnet) course I already have experience with Firestore and XAMPP SQL database impelemtations. Therefore, I wanted to test MongoDB in this project.

## Planned implementations
<ul>
  <li>Simple frontend home page, and other required pages for navigation. Low priority.</li>
  <li>Weather report for a given location.</li>
  <li>A recommendation system based on a given movie, though the selection of recommended movies will be limited by the number of movies included in the generated MongoDB sample collection.</li>
  <li>Football statistics for a given team.</li>
  <li>The service can register Webhooks for notifications for events happening with their football team</li>
</ul>

# Endpoints

## Registrations

### POST
The `POST` request to the registrations endpoint allows users to register an account. 
The request body must contain the `username`, `password` and `email` fields. 
The `password` field must contain all numbers to ensure no proper password is used in this dummy service.
The password is hashed before it is stored in the database.
<br>
The request body also contains boolean values for what the user wishes to have available on their dashboards. The `preferences` field is optional, and if omitted, the user will have no preferences set.

**Example JSON body:**
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

// TODO
