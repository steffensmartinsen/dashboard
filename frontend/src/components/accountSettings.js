import Header from "./header";
import { useNavigate } from "react-router-dom";

const EditAccount = (props) => {
    const navigate = useNavigate();

    console.log("EditAccount.js: loggedIn: ", props.loggedIn);

    if (!props.loggedIn) {
        navigate('/login');
    }

    return (
        <div className={'mainContainer'}>
            <Header username={props.username} loggedIn={props.loggedIn} setLoggedIn={props.setLoggedIn} />
            <div className={'titleContainer'}>
                <div>Edit Account</div>
            </div>
        </div>
    )
}

export default EditAccount;
