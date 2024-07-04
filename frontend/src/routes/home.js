import React from 'react'
import { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import Header from "../components/header";
import { Button } from '@chakra-ui/react'
import { GetCookie, Logout } from "../utils/helpers";
import  WeatherSquare  from "../components/weatherSquare";

const Home = (props) => {
    const { loggedIn, username,  setLoggedIn } = props
    const [token, setToken] = useState('')
    const [loading, setLoading] = useState(true)
    const navigate = useNavigate()


    // Helper function to render the content based on the loggedIn state
    function renderContent() {
        if (loggedIn) {
            return (
                <>
                    < WeatherSquare username={username} />
                    <div className='titleContainer'>
                        Welcome, {username}!
                    </div>
                </>
            );
        } else {
            return (
                <div className='titleContainer'>
                    Welcome!
                </div>
            );
        }
    }


    useEffect(() => {
        let storedUsername = ""
        if (username === '') {
            storedUsername = localStorage.getItem('username') || '';
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
            navigate('/login')
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
                {renderContent()}
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