import "../App.css";
import { ImGithub } from "react-icons/im";
import { TiHome } from "react-icons/ti";
import { FaInfoCircle } from "react-icons/fa";

const Header = () => {

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
                        <a href="https://github.com/viatheboy/dashboard" target="_blank">
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
        </header>
    )
}

export default Header;