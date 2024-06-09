import React, {useState} from "react";
import "../App.css";
import Header from "./header";
import UsernameInput from "./usernameInput";
import PasswordInput from "./passwordInput";
import EmailInput from "./emailInput";
import EnforcePassword from "../utils/helpers";
import {Button} from "@chakra-ui/react";
import enforcePassword from "../utils/helpers";

const Register = (props) => {
    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const [repeatedPassword, setRepeatedPassword] = useState('')
    const [email, setEmail] = useState('')
    const [errorMessage, setErrorMessage] = useState('')
    // const [emailError, setEmailError] = useState('')
    // const [usernameError, setUsernameError] = useState('')
    // const [passwordError, setPasswordError] = useState('')

    const onButtonClick = () => {
        // Set initial error value to empty
        setErrorMessage('')

        // Check if the user has entered all fields correctly
        if (username === '') {
            setErrorMessage('Username is required')
            return
        }
        if (username != username.toLowerCase()) {
            setErrorMessage('Username must be lowercase')
            return
        }
        if (email === '') {
            setErrorMessage('Email is required')
            return
        }
        if (!/^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/.test(email)) {
            setErrorMessage('Please enter a valid email address')
            return
        }
        if (password === '') {
            setErrorMessage('Password is required')
            return
        }
        if (password.length < 8) {
            setErrorMessage('Password must be at least 8 characters')
            return
        }
        if (password !== repeatedPassword) {
            setErrorMessage('Passwords do not match')
            return
        }
        if (!enforcePassword(password)) {
            setErrorMessage("Wrong password. " +
                "Please don't use an actual password for this. " +
                "The only accepted characters are '1234567890'")
        }


    }

    return (
        <div className={'mainContainer'}>
            <Header/>
            <div className={'titleContainer'}>
                <div>Register</div>
            </div>
            <div className={'inputContainer'}>
                <UsernameInput
                    username={username}
                    setUsername={setUsername}
                    onChange={(ev) => setUsername(ev.target.value)}
                    className='registerInput'
                />
            </div>
            <div className="inputContainer">
                <EmailInput
                    email={email}
                    onChange={(ev) => setEmail(ev.target.value)}
                    className='registerInput'
                />
            </div>
            <div className={'inputContainer'}>
                <PasswordInput
                    password={password}
                    onChange={(ev) => setPassword(ev.target.value)}
                    className='registerInput'
                    placeholder='Enter password'
                />
            </div>
            <div className={'inputContainer'}>
                <PasswordInput
                    password={repeatedPassword}
                    onChange={(ev) => setRepeatedPassword(ev.target.value)}
                    className='registerInput'
                    placeholder='Repeat password'
                />
            </div>
            <div className='inputContainer'>
                <label className="errorLabel">{errorMessage}</label>
            </div>
            <div className={'buttonContainer'}>
                <Button colorScheme='teal' size='md' onClick={onButtonClick} className={'loginButton'}>
                    Register
                </Button>
            </div>
        </div>
    )
}
export default Register