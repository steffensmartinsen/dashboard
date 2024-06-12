import Header from "./header";
import { useNavigate } from "react-router-dom";
import {Button, FormControl, FormLabel, Input, InputGroup, Switch} from "@chakra-ui/react";
import { React, useEffect, useState } from "react";
import {GetCookie, GetUser} from "../utils/helpers";
import EmailInput from "./emailInput";
import FootballInput from "./footballInput";

const EditAccount = (props) => {

    // Declare variables
    const [loading, setLoading] = useState(true);
    const {username, setLoggedIn, loggedIn} = props;
    const navigate = useNavigate();
    const [email, setEmail] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    const [user, setUser] = useState({});
    const [weather, setWeather] = useState(false);
    const [football, setFootball] = useState(false);
    const [movies, setMovies] = useState(false);
    const [team, setTeam] = useState('');

    // Check if the user is logged in, redirect if not
    useEffect(() => {
        let storedUsername = ""
        if (username === '') {
            storedUsername = localStorage.getItem('username') || '';
            props.setUsername(storedUsername)
        } else {
            storedUsername = username;
        }
        if (storedUsername) {
            GetCookie(storedUsername, setLoggedIn).then(() => {
                if (localStorage.getItem('loggedIn') === 'false') {
                    navigate('/login');
                }
            });
        } else {
            if (localStorage.getItem('loggedIn') === 'false') {
                navigate('/login');
            }
        }
    }, [username, setLoggedIn]);

    useEffect(() => {
        GetUser(username, (data) => {
            setUser(data);
            console.log(data);
            setEmail(data.email);
            setWeather(data.preferences.weather);
            setFootball(data.preferences.football);
            setMovies(data.preferences.movies);
            setTeam(data.preferences.team);
            setLoading(false)
        }).catch((error) => {
            console.error('Error:', error)
        });
    }, [username]);


    // Render the page
    if (loading) {
        return (
            <div className={'mainContainer'}>
                <Header username={username} loggedIn={loggedIn} setLoggedIn={setLoggedIn}/>
            </div>
        )
    }

    console.log(weather);

    return (
        <div className={'mainContainer'}>
            <Header username={username} loggedIn={loggedIn} setLoggedIn={setLoggedIn}/>
            <div className={'titleContainer'}>
                <div>Account Details</div>
            </div>
            <div className={'inputContainer'}>
                <InputGroup size='md' className={"usernameNoAccess"}>
                    <Input
                        pr='4.5rem'
                        type={"text"}
                        placeholder={username}
                        isDisabled={true}
                    />
                </InputGroup>
            </div>
            <div className={'inputContainer'}>
                <EmailInput
                    email={email}
                    onChange={(ev) => setEmail(ev.target.value)}
                    className='registerInput'
                />
            </div>
            <div className='inputContainer'>
                <label className="errorLabel">{errorMessage}</label>
            </div>
            <div className='subtitleContainer'>
                Preferences
            </div>
            <div className='inputContainer'>
                <FormControl display='flex' alignItems='center' className='switchForm'>
                    <FormLabel mb='0'>
                        <div className='switchLabel'>Weather Report?</div>
                    </FormLabel>
                    <Switch
                        colorScheme='teal'
                        defaultChecked={weather}
                        onChange={() => setWeather(!weather)}
                    />
                </FormControl>
            </div>
            <div className='inputContainer'>
                <FormControl display='flex' alignItems='center' className='switchForm'>
                    <FormLabel mb='0' >
                        <div className='switchLabel'>Football Updates?</div>
                    </FormLabel>
                    <Switch
                        colorScheme='teal'
                        defaultChecked={football}
                        onChange={() => setFootball(!football)}
                    />
                </FormControl>
            </div>
            {football && (
                <div className='inputContainer'>
                    <FootballInput team={team} onChange={(ev) => setTeam(ev.target.value)} className='registerInput'/>
                </div>
            )}
            <div className='inputContainer'>
                <FormControl display='flex' alignItems='center' className='switchForm'>
                    <FormLabel mb='0'>
                        <div className='switchLabel'>Movie Recommendations?</div>
                    </FormLabel>
                    <Switch
                        colorScheme='teal'
                        defaultChecked={movies}
                        onChange={() => setMovies(!movies)}
                    />
                </FormControl>
            </div>
            <div className={'registerButtonContainer'}>
                <div>Button</div>
            </div>
        </div>
    )
}

export default EditAccount;
