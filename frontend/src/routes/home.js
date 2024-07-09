import React from 'react'
import { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import Header from "../components/header";
import { Button } from '@chakra-ui/react'
import { GetCookie, Logout } from "../utils/helpers";
import { RenderMainContent } from '../utils/helpers';
import { LOGIN, USERNAME } from '../utils/consts';

const Home = (props) => {
    const { loggedIn, username,  setLoggedIn } = props
    const [token, setToken] = useState('')
    const [loading, setLoading] = useState(true)
    const navigate = useNavigate()

    useEffect(() => {
        let storedUsername = ""
        if (username === '') {
            storedUsername = localStorage.getItem(USERNAME) || '';
            props.setUsername(storedUsername)
        }
        if (storedUsername) {
            GetCookie(storedUsername, setLoggedIn).then(() => {
                setLoading(false)
            });
        } else {
            setLoading(false);
        }
    }, [username, setLoggedIn]);

    const onButtonClick = () => {
        if (loggedIn) {
            Logout(username, setLoggedIn)
        } else {
            navigate(LOGIN)
        }
    }

    if (loading) {
        return (
            <div className="mainContainer">
                <Header username={username} loggedIn={loggedIn} setLoggedIn={setLoggedIn} />
            </div>
        )
    }

    return (
        <div className='mainContainer'>
            <Header username={username} loggedIn={loggedIn} setLoggedIn={setLoggedIn} />
            <div className='mainContent'>
                {RenderMainContent(loggedIn, username)}
                <div className='buttonContainer'>
                    <Button colorScheme='teal' size='md' onClick={onButtonClick} className='loginButton'>
                        {loggedIn ? 'Log out' : 'Log in'}
                    </Button>
                </div>
            </div>
        </div>
    )
}

export default Home