import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom'

const Login = (props) => {
    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const [usernameError, setUsernameError] = useState('')
    const [passwordError, setPasswordError] = useState('')

    const navigate = useNavigate()

    const onButtonClick = () => {
        // Set initial error values to empty
        setUsernameError('')
        setPasswordError('')

        if (username === '') {
            setUsernameError("Username is required")
            return
        }

        if (password === '') {
            setPasswordError("Password is required")
            return
        }

        if (password.length < 8) {
            setPasswordError("Password must be at least 8 characters")
            return
        }

        checkAccountExists((response, status) => {
            if (status === 200) {
                window.alert("Account with username " + username + " found in the database")
            } else if (status === 404){
                window.alert("Account with username " + username + " not found in the database")
            } else {
                window.alert("An error occured. Status code: " + status)
            }
        }, username)
    }

    // Call the server API to check if the given email ID already exists
    const checkAccountExists = (callback, username) => {
        fetch("http://localhost:8080/dashboards/v1/registrations/" + username + "/", {
            method: 'GET',
        })
            .then((r) => {
                // Only parse JSON if the status code is 200
                if (r.status === 200) {
                    r.json().then((json) => callback(json, r.status))
                } else {
                    callback(null, r.status)
                }
            })
            .catch((error) => {
                console.error('Error:', error)
            })
        }

    return (
        <div className={'mainContainer'}>
            <div className={'titleContainer'}>
                <div>Login</div>
            </div>
            <br />
            <div className={'inputContainer'}>
                <input
                    value={username}
                    placeholder="Enter your username here"
                    onChange={(ev) => setUsername(ev.target.value)}
                    className={'inputBox'}
                />
                <label className="errorLabel">{usernameError}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <input
                    type="password"
                    value={password}
                    placeholder="Enter your password here"
                    onChange={(ev) => setPassword(ev.target.value)}
                    className={'inputBox'}
                />
                <label className="errorLabel">{passwordError}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <input className={'inputButton'} type="button" onClick={onButtonClick} value={'Log in'} />
            </div>
        </div>
    )
}

export default Login