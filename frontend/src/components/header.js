import "../App.css";
import { ImGithub } from "react-icons/im";
import { TiHome } from "react-icons/ti";
import { FaInfoCircle } from "react-icons/fa";
import { useNavigate } from "react-router-dom";
import {
    Menu,
    MenuButton,
    MenuList,
    MenuItem,
    Button,
    IconButton,
} from '@chakra-ui/react'
import { HamburgerIcon } from '@chakra-ui/icons';

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
                {!loggedIn ? (
                <div className={'headerRight'}>
                    <Button colorScheme="teal" variant="solid" size="sm" className={'headerButton'}
                            onClick={() => navigate('/login')}>
                        Log in
                    </Button>
                    <Button colorScheme="teal" variant="solid" size="sm"  className={'headerButton'}
                            onClick={() => navigate('/register')}>
                        Sign up
                    </Button>
                    </div>
                    ) : (
                        <div className={'headerRight'}>
                            <Menu>
                                <MenuButton
                                    as={IconButton}
                                    aria-label='Options'
                                    icon={<HamburgerIcon />}
                                    variant='outline'
                                    color={'white'}
                                    backgroundColor={'hsl(180, 100%, 20%)'}
                                />
                                <MenuList>
                                    <MenuItem color={'teal' } fontWeight={'bold'}>
                                        Account Settings
                                    </MenuItem>
                                    <MenuItem color={'teal'} fontWeight={'bold'} >
                                        Log out
                                    </MenuItem>
                                </MenuList>
                            </Menu>
                        </div>
                    )}

        </header>
    )
}

export default Header;