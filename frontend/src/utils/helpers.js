
// Function to enforce only numerical characters for passwords
function EnforcePassword(password) {
    const accepted = ["1", "2", "3", "4", "5", "6", "7", "8", "9", "0"];

    for (let char of password) {
        let found = false;
        for (let element of accepted) {
            if (char === element) {
                found = true;
                break;
            }
        }
        if (!found) {
            return false;
        }
    }
    return true;
}

// Function to create a new user through backend API
const CreateUser = (callback, data) => {
    fetch("http://localhost:8080/dashboards/v1/registrations/", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    })
        .then((r) => {
            callback(r.status)
        })
        .catch((error) => {
            console.error('Error:', error)
        })
}

// Function to authenticate a user through backend API
const AuthenticateUser = async (callback, username, password) => {
    fetch("http://localhost:8080/dashboards/v1/auth/", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            username,
            password,
        }),
    })
        .then((r) => {
            callback(r.status)
        })
        .catch((error) => {
            console.error('Error:', error)
        })
}

const SetCookie =  async (username) => {
    try {
        const response = await fetch("http://localhost:8080/dashboards/v1/set-cookie/", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ username }),
            credentials: 'include',
        });

        if (response.ok) {
            console.log("Cookie set successfully");
            return true;
        } else {
            console.log("Failed to set cookie")
            return false;
        }
    } catch (error) {
        console.error('Error:', error)
        return false;
    }
}

const GetCookie = async (username, setToken) => {
    try {
        const response = await fetch("http://localhost:8080/dashboards/v1/get-cookie/" + username + "/", {
            method: 'GET',
            credentials: 'include',
        });

        if (response.ok) {
            const token = await response.text()
            setToken(token)
        } else {
            console.log("Failed to get cookie")
        }
    } catch (error) {
        console.error('Error:', error)
    }
}

export { EnforcePassword, CreateUser, AuthenticateUser, SetCookie, GetCookie };