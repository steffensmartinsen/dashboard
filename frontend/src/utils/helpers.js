
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

// SetCookie function to set a cookie in the browser and adds the username to local storage
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
            localStorage.setItem('username', username)
        } else {
            console.log("Failed to set cookie")
        }
    } catch (error) {
        console.error('Error:', error)
    }
}

const GetCookie = async (username, setLoggedIn) => {
    try {
        const response = await fetch("http://localhost:8080/dashboards/v1/get-cookie/" + username + "/", {
            method: 'GET',
            credentials: 'include',
        });

        if (response.ok) {
            const token = await response.text()
            console.log("TOKEN: ", token);
            setLoggedIn(true)
        } else {
            console.log("Failed to get cookie")
            setLoggedIn(false)
        }
    } catch (error) {
        console.error('Error:', error)
    }
}

const DeleteCookie = (username) => {
    try {
        const response = fetch("http://localhost:8080/dashboards/v1/delete-cookie/" + username + "/", {
            method: 'DELETE',
            credentials: 'include',
        });

        if (response.ok) {
            console.log("Cookie deleted successfully");
            localStorage.clear()
        } else {
            console.log("Failed to delete cookie")
        }
    } catch (error) {
        console.error('Error:', error)
    }
}

export { EnforcePassword, CreateUser, AuthenticateUser, SetCookie, GetCookie, DeleteCookie };