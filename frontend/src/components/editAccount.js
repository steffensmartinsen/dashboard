import Header from "./header";
import { useNavigate } from "react-router-dom";
import {Input, InputGroup} from "@chakra-ui/react";
import { React, useEffect, useState } from "react";
import {GetCookie} from "../utils/helpers";

const EditAccount = (props) => {
    const [loading, setLoading] = useState(true);
    const {username, setLoggedIn, loggedIn} = props;
    const navigate = useNavigate();

    console.log("EditAccount.js: loggedIn: ", props.loggedIn);

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
                setLoading(false)
                if (localStorage.getItem('loggedIn') === 'false') {
                    navigate('/login');
                }
            });
        } else {
            setLoading(false);
            if (localStorage.getItem('loggedIn') === 'false') {
                navigate('/login');
            }
        }
    }, [username, setLoggedIn]);

    if (loading) {
        return (
            <div className={'mainContainer'}>
                <Header username={username} loggedIn={loggedIn} setLoggedIn={setLoggedIn}/>
            </div>
        )
    }

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
        </div>
    )
}

export default EditAccount;
