
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

// GetUser function to get a user from the backend API
const GetUser = async (username, callback) => {
    try {
        const response = await fetch("http://localhost:8080/dashboards/v1/registrations/" + username + "/", {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
        });

        if (response.ok) {
            const data = await response.json();
            callback(data)
        } else {
            console.log("Failed to get user")
        }
    } catch (error) {
        console.error('Error:', error)
    }
}

// CreateUser function to create a new user through backend API
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

const UpdateUser = (callback, data) => {
    fetch("http://localhost:8080/dashboards/v1/registrations/", {
        method: 'PUT',
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

// AuthenticateUser function to authenticate a user through backend API
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

// GetCookie function to get a cookie from the browser and sets the loggedIn state
const GetCookie = async (username, setLoggedIn) => {
    try {
        const response = await fetch("http://localhost:8080/dashboards/v1/get-cookie/" + username + "/", {
            method: 'GET',
            credentials: 'include',
        });

        if (response.ok) {
            console.log("Cookie retrieved successfully");
            localStorage.setItem('loggedIn', 'true')
            setLoggedIn(true)
        } else {
            console.log("Failed to get cookie")
            localStorage.setItem('loggedIn', 'false')
            setLoggedIn(false)
        }
    } catch (error) {
        console.error('Error:', error)
    }
}

// DeleteCookie function to delete a cookie from the browser and clear local storage
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

// Logout function to delete the cookie and clear local storage
const Logout = (username, setLoggedIn) => {
    DeleteCookie(username);
    localStorage.clear();
    setLoggedIn(false);
}

// UsernameCheck function to ensure username is correct
const UsernameCheck = (username, setErrorMessage) => {
    if (username === '') {
        setErrorMessage('Username is required')
        return false
    }
    if (username !== username.toLowerCase()) {
        setErrorMessage('Username must be lowercase')
        return false
    }
    return true
}

// EmailCheck function to ensure email is correct
const EmailCheck = (email, setErrorMessage) => {
    if (email === '') {
        setErrorMessage('Email is required')
        return false
    }
    if (!/^[\w-\.øæå]+@([\w-\.øæå]+\.)+[\w-\.øæå]{2,4}$/.test(email)) {
        setErrorMessage('Please enter a valid email address')
        return false
    }
    return true
}

// PasswordCheck function to ensure password is correct and matches the repeated password
const PasswordCheck = (password, repeatedPassword, setErrorMessage) => {
    if (password === '') {
        setErrorMessage('Password is required')
        return false
    }
    if (password.length < 8) {
        setErrorMessage('Password must be at least 8 characters')
        return false
    }
    if (password !== repeatedPassword) {
        setErrorMessage('Passwords do not match')
        return false
    }
    if (!EnforcePassword(password)) {
        setErrorMessage("Wrong password. " +
            "Please don't use an actual password for this. " +
            "The only accepted characters are '1234567890'")
        return false
    }
    return true;
}

export { EnforcePassword, CreateUser, GetUser, AuthenticateUser, SetCookie, GetCookie, DeleteCookie, Logout, PasswordCheck, EmailCheck, UsernameCheck, UpdateUser };