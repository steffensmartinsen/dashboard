import Header from "./header";
import {React} from "react";

const ChangePassword = (props) => {
    const { loggedIn, username, setLoggedIn } = props;

    return (
        <div className={'mainContainer'}>
            <Header username={username} loggedIn={loggedIn} setLoggedIn={setLoggedIn}/>
            <div className={'titleContainer'}>
                <div>Change Password</div>
            </div>
        </div>
    )
}

export default ChangePassword;