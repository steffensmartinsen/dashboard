import React, {useState} from "react";
import "../App.css";
import Header from "../components/header";
import UsernameInput from "../components/usernameInput";
import PasswordInput from "../components/passwordInput";
import EmailInput from "../components/emailInput";
import FootballInput from "../components/footballInput";
import {useNavigate} from "react-router-dom";
import {Button, Switch, FormControl, FormLabel, InputGroup, Input} from "@chakra-ui/react";
import { PasswordCheck, CreateUser, SetCookie, EmailCheck, UsernameCheck, CountryAndCityCheck } from "../utils/helpers";
import CountrySelector from "../components/countrySelector";
import { ERROR_500, ERROR_EXISTS, ERROR_SELECT_TEAM, ERROR_UNDEFINED, ROOT, SUCCESS_USER_CREATED } from "../utils/consts";

// Component to create a new user through backend API
const Register = (props) => {
    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const [repeatedPassword, setRepeatedPassword] = useState('')
    const [email, setEmail] = useState('')
    const [errorMessage, setErrorMessage] = useState('')
    const [footballError, setFootballError] = useState("")
    const [country, setCountry] = useState('')
    const [city, setCity] = useState('')
    const [weather, setWeather] = useState(false)
    const [football, setFootball] = useState(false)
    const [movie, setMovie] = useState(false)
    const [team, setTeam] = useState('')


    const navigate = useNavigate()

    // Function to handle the button click
    const onButtonClick = () => {

        // Set initial error value to empty
        setErrorMessage('')

        // Run checks on username, email and password
        if (!UsernameCheck(username, setErrorMessage) || !EmailCheck(email, setErrorMessage) || !PasswordCheck(password, repeatedPassword, setErrorMessage)) {
            return
        }

        // Check if the user has selected a country and a city
        if (!CountryAndCityCheck(country, city, setErrorMessage)) {
            return
        }

        if (football && team === "") {
            setFootballError(ERROR_SELECT_TEAM)
            return
        }

        // Create data object to send to backend
        const data = {
            "username": username,
            "email": email,
            "password": password,
            "country": country,
            "city": city,
            "preferences": {
                "football": football,
                "weather": weather,
                "movies": movie,
                "team": team,
            },
        }

        // Call to CreateUser function
        CreateUser((status) => {
            switch (status) {
                case 201:
                    console.log(SUCCESS_USER_CREATED);
                    props.setLoggedIn(true);
                    props.setUsername(username);
                    SetCookie(username);
                    navigate(ROOT)
                    break;
                case 400:
                    setErrorMessage(ERROR_EXISTS);
                    break;
                case 500:
                    setErrorMessage(ERROR_500);
                    break;
                default:
                    setErrorMessage(ERROR_UNDEFINED);
            }
        }, data)
    }

    return (
        <div className={'mainContainer'}>
            <Header/>
            <div className='mainContent'>
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
                        autoComplete='off'
                    />
                </div>
                <div className={"inputContainer"}>
                    <CountrySelector country={country} setCountry={setCountry} className='registerInput'/>
                </div>
                <div className={'inputContainer'}>
                    <InputGroup size='md' className={"registerInput"}>
                        <Input
                            pr='4.5rem'
                            type={"text"}
                            onChange={(ev) => setCity(ev.target.value)}
                            placeholder={"Enter your city"}
                        />
                    </InputGroup>
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
                            <div className='switchLabel'>Weather Report</div>
                        </FormLabel>
                        <Switch
                            colorScheme='teal'
                            onChange={() => setWeather(!weather)}
                        />
                    </FormControl>
                </div>
                <div className='inputContainer'>
                    <FormControl display='flex' alignItems='center' className='switchForm'>
                        <FormLabel mb='0'>
                            <div className='switchLabel'>Football Updates</div>
                        </FormLabel>
                        <Switch
                            colorScheme='teal'
                            onChange={() => setFootball(!football)}
                        />
                    </FormControl>
                </div>
                {football && (
                    <div className='inputContainer'>
                        <FootballInput team={team} onChange={(ev) => setTeam(ev.target.value)} className='registerInput'/>
                    </div>
                )}
                {football && (
                    <div className='inputContainer'>
                        <label className="errorLabel">{footballError}</label>
                    </div>
                )}
                <div className='inputContainer'>
                    <FormControl display='flex' alignItems='center' className='switchForm'>
                        <FormLabel mb='0'>
                            <div className='switchLabel'>Movie Recommendations</div>
                        </FormLabel>
                        <Switch
                            colorScheme='teal'
                            onChange={() => setMovie(!movie)}
                        />
                    </FormControl>
                </div>
                <div className={'registerButtonContainer'}>
                    <Button colorScheme='teal' size='md' onClick={onButtonClick} className={'loginButton'}>
                        Sign up
                    </Button>
                </div>
            </div>
        </div>
    )
}
export default Register