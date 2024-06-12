import React from 'react'
import { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import Header from "./header";
import { Button } from '@chakra-ui/react'
import { GetCookie, DeleteCookie } from "../utils/helpers";

const Home = (props) => {
    const { loggedIn, username,  setLoggedIn } = props
    const [token, setToken] = useState('')
    const [loading, setLoading] = useState(true)
    const navigate = useNavigate()


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

    console.log("Home.js: loggedIn: ", loggedIn)

    const onButtonClick = () => {
        if (loggedIn) {
            props.setLoggedIn(false)
            DeleteCookie(username, setLoggedIn, setLoggedIn)
            localStorage.clear()
        } else {
            navigate('/login')
        }
    }

    if (loading) {
        return (
            <div className="mainContainer">
                <Header loggedIn={loggedIn} setLoggedIn={setLoggedIn} />
            </div>
        )
    }

    return (
        <div className="mainContainer">
            <Header loggedIn={loggedIn} setLoggedIn={setLoggedIn} />
            <div className={'titleContainer'}>
                {loggedIn ? <div>Welcome, {username}!</div> : <div>Welcome!</div>}
            </div>
            <div>This is the home page.</div>
            <div className={'buttonContainer'}>
                <Button colorScheme="teal" size="md" onClick={onButtonClick} className={'loginButton'}>
                    {loggedIn ? 'Log out' : 'Log in'}
                </Button>
            </div>
        </div>
    )
}

export default Home