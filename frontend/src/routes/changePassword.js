import Header from "../components/header";
import {React, useEffect, useState} from "react";
import PasswordInput from "../components/passwordInput";
import {useNavigate} from "react-router-dom";
import {GetCookie, PasswordCheck, AuthenticateUser, GetUser, UpdateUser} from "../utils/helpers";
import {Button} from "@chakra-ui/react";
import { LOGGEDIN, LOGIN, ROOT, VALUE_FALSE } from "../utils/consts";

const ChangePassword = (props) => {

    const [loading, setLoading] = useState(true);
    const { loggedIn, username, setLoggedIn } = props;
    const [currentPassword, setCurrentPassword] = useState('');
    const [newPassword, setNewPassword] = useState('');
    const [repeatedPassword, setRepeatedPassword] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    const [passwordMessage, setPasswordMessage] = useState('')
    const [authenticated, setAuthenticated] = useState(false);
    const navigate = useNavigate();

    // Check if the user is logged in, redirect if not
    useEffect(() => {
        let storedUsername = ""
        if (username === '') {
            storedUsername = localStorage.getItem(username) || '';
            props.setUsername(storedUsername)
        } else {
            storedUsername = username;
        }
        if (storedUsername) {
            GetCookie(storedUsername, setLoggedIn).then(() => {
                if (localStorage.getItem(LOGGEDIN) === VALUE_FALSE) {
                    navigate(LOGIN);
                }
            });
        } else {
            if (localStorage.getItem(LOGGEDIN) === VALUE_FALSE) {
                navigate(LOGIN);
            }
        }
    }, [username, setLoggedIn]);

    // Render the page
    if (loading) {
        return (
            <div className={'mainContainer'}>
                <Header username={username} loggedIn={loggedIn} setLoggedIn={setLoggedIn}/>
            </div>
        )
    }

    const onButtonClick = async () => {

        // Set initial error value to empty
        setErrorMessage('')

        // Check the new password
        if (!PasswordCheck(newPassword, repeatedPassword, setErrorMessage)) {
            return
        }

        // Call to AuthenticateUser function to check the current password
        await AuthenticateUser(status => {
            if (status === 200) {
                setAuthenticated(true)
            } else {
                setPasswordMessage("Incorrect password")
                return
            }
        }, username, currentPassword)

        let userData = null;
        // Call to GetUser function to get the user data
        await GetUser(username, (data) => {
            userData = data;
        }).catch((error) => {
            console.error('Error:', error)
        });

        // Change the password field
        userData.password = newPassword;

        // Call to UpdateUser function to update the user data
        UpdateUser((status) => {
            switch (status) {
                case 200:
                    console.log('Password updated successfully');
                    navigate(ROOT)
                    break;
                default:
                    setErrorMessage('Something went wrong');
            }
        }, userData)


    }


    return (
        <div className={'mainContainer'}>
            <Header username={username} loggedIn={loggedIn} setLoggedIn={setLoggedIn}/>
            <div className={'titleContainer'}>
                <div>Change Password</div>
            </div>
            <div className={'inputContainer'}>
                <div className={"inputLabel"}>Current Password:</div>
                <PasswordInput
                    password={currentPassword}
                    onChange={(ev) => setCurrentPassword(ev.target.value)}
                    className='registerInput'
                    placeholder='Current password'
                    autocomplete='off'
                />
            </div>
            <div className='inputContainer'>
                <label className="errorLabel">{passwordMessage}</label>
            </div>
            <div className={'inputContainer'}>
                <div className={"inputLabel"}>New Password:</div>
                <PasswordInput
                    password={newPassword}
                    onChange={(ev) => setNewPassword(ev.target.value)}
                    className='registerInput'
                    placeholder='New password'
                    autocomplete='off'
                />
            </div>
            <div className={'inputContainer'}>
                <div className={"inputLabel"}>Repeat New Password:</div>
                <PasswordInput
                    password={repeatedPassword}
                    onChange={(ev) => setRepeatedPassword(ev.target.value)}
                    className='registerInput'
                    placeholder='Repeat new password'
                    autoComplete='off'
                />
            </div>
            <div className='inputContainer'>
                <label className="errorLabel">{errorMessage}</label>
            </div>
            <div className={'registerButtonContainer'}>
                <Button colorScheme='teal' size='md' onClick={onButtonClick} className={'loginButton'}>
                    Save
                </Button>
            </div>
        </div>
    )
}

export default ChangePassword;