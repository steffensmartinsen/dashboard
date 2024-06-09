import {useState} from "react";
import Header from "./header";

const Register = (props) => {
    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const [usernameError, setUsernameError] = useState('')
    const [passwordError, setPasswordError] = useState('')

    const onButtonClick = () => {
        //TODO
    }

    return (
        <div className={'mainContainer'}>
            <Header />
            <div className={'titleContainer'}>
                <div>Register</div>
            </div>
            <br/>
            <div className={'inputContainer'}>
                <input
                    value={username}
                    placeholder="Enter your username here"
                    onChange={(ev) => setUsername(ev.target.value)}
                    className={'inputBox'}
                />
                <label className="errorLabel">{usernameError}</label>
            </div>
            <br/>
            <div className={'inputContainer'}>
                <input
                    value={password}
                    placeholder="Enter your password here"
                    onChange={(ev) => setPassword(ev.target.value)}
                    className={'inputBox'}
                />
                <label className="errorLabel">{passwordError}</label>
            </div>
            <br/>
            <div className={'buttonContainer'}>
                <input
                    className={'inputButton'}
                    type="button"
                    onClick={onButtonClick}
                    value={'Register'}
                />
            </div>
        </div>
    )
}
export default Register