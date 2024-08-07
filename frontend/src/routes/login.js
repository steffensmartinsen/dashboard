import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import PasswordInput from "../components/passwordInput";
import UsernameInput from "../components/usernameInput";
import Header from "../components/header";
import { Button } from '@chakra-ui/react'
import { AuthenticateUser, SetCookie } from '../utils/helpers'
import { ERROR_LOGIN, ERROR_PASSWORD_REQUIRED, ERROR_USERNAME_REQUIRED, ROOT, SUCCESS_USER_AUTHENTICATED } from '../utils/consts';


const Login = (props) => {
    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const [usernameError, setUsernameError] = useState('')
    const [passwordError, setPasswordError] = useState('')

    const navigate = useNavigate()

    const onButtonClick = async () => {
        // Set initial error values to empty
        setUsernameError('')
        setPasswordError('')

        if (username === '') {
            setUsernameError(ERROR_USERNAME_REQUIRED)
            return
        }

        if (password === '') {
            setPasswordError(ERROR_PASSWORD_REQUIRED)
            return
        }

        await AuthenticateUser(status => {
            if (status === 200) {
                console.log(SUCCESS_USER_AUTHENTICATED)
                props.setLoggedIn(true)
                props.setUsername(username)
                SetCookie(username)
                navigate(ROOT)
            } else {
                setPasswordError(ERROR_LOGIN)
                return
            }
        }, username, password)

    }

    return (
        <div className={'mainContainer'}>
            <Header />
            <div className="mainContent">
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
                        autoComplete='on'
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
        </div>
    )
}

export default Login