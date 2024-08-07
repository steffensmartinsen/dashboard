import "../App.css";
import { ImGithub } from "react-icons/im";
import { TiHome } from "react-icons/ti";
import { FaInfoCircle } from "react-icons/fa";
import { useNavigate } from "react-router-dom";
import { Button } from '@chakra-ui/react'
import BurgerMenu from "./burgerMenu";
import { EXTERNAL_GITHUB, LOGIN, REGISTER } from "../utils/consts";

const Header = (props) => {
    const { loggedIn, setLoggedIn } = props;
    const navigate = useNavigate();

    return (
        <header className="header">
                <nav>
                    <ul>
                        <li>
                            <a href="/">
                                <div className="iconText">
                                    <TiHome />
                                    <span className='iconSpan'>Home</span>
                                </div>
                            </a>
                        </li>
                        <li>|</li>
                        <li>
                            <a href={EXTERNAL_GITHUB} target="_blank">
                                <div className="iconText">
                                    <ImGithub />
                                    <span className='iconSpan'>Github</span>
                                </div>
                            </a>
                        </li>
                        <li>|</li>
                        <li>
                            <a href="/">
                                <div className="iconText">
                                    <FaInfoCircle />
                                    <span className='iconSpan'>About</span>
                                </div>
                            </a>
                        </li>
                    </ul>
                </nav>
                {!loggedIn ? (
                <div className={'headerRight'}>
                    <Button colorScheme="teal" variant="solid" size="sm" className={'headerButton'}
                            onClick={() => navigate(LOGIN)}>
                        Log in
                    </Button>
                    <Button colorScheme="teal" variant="solid" size="sm"  className={'headerButton'}
                            onClick={() => navigate(REGISTER)}>
                        Sign up
                    </Button>
                    </div>
                    ) : (
                        <div className={'headerRight'}>
                            <BurgerMenu username={props.username} setLoggedIn={setLoggedIn}/>
                        </div>
                    )}

        </header>
    )
}

export default Header;