import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import PasswordInput from "./passwordInput";
import UsernameInput from "./usernameInput";
import Header from "./header";
import { Button } from '@chakra-ui/react'
import { AuthenticateUser } from '../utils/helpers'


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

        AuthenticateUser(status => {
            if (status === 200) {
                console.log("User authentication successful")
                props.setLoggedIn(true)
                props.setUsername(username)
                navigate('/')
            } else {
                setPasswordError("Invalid username or password")
            }
        }, username, password)
    }

    return (
        <div className={'mainContainer'}>
            <Header />
            <div className={'titleContainer'}>
                <div>Login</div>
            </div>
            <br />
            <div className={'inputContainer'}>
                <UsernameInput
                    username={username}
                    onChange={(ev) => setUsername(ev.target.value)}
                    className={'registerInput'}
                />
                <label className="errorLabel">{usernameError}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <PasswordInput
                    password={password}
                    onChange={(ev) => setPassword(ev.target.value)}
                    placeholder='Enter password'
                    className={'registerInput'}
                />
                <label className="errorLabel">{passwordError}</label>
            </div>
            <br />
            <div className={'inputContainer'}>
                <Button colorScheme='teal' size='md' onClick={onButtonClick} className={'loginButton'}>
                    Log in
                </Button>
            </div>
        </div>
    )
}

export default Login