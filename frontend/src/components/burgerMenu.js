import { useState } from "react";
import { GiHamburgerMenu } from "react-icons/gi";
import { IoMdClose } from "react-icons/io";
import { Logout } from "../utils/helpers";
import { useNavigate } from "react-router-dom";

const BurgerMenu = (props) => {
    const [menuOpen, setMenuOpen] = useState(false);
    const navigate = useNavigate();

    const handleMenuClick = () => {
        setMenuOpen(!menuOpen);
    }

    // Function to handle logout
    const onClickLogout = () => {
        Logout(props.username, props.setLoggedIn);
        navigate('/');
    }

    return (
        <div className="burgerMenu">
            <div className="burgerMenuIcon" onClick={handleMenuClick}>
                <GiHamburgerMenu />
            </div>
            {menuOpen && (
                <div className="burgerMenuContent">
                    <div className="closeIcon" onClick={handleMenuClick}>
                        <IoMdClose/>
                    </div>
                    <div className="link" onClick={() => navigate("/account")}>Account Details</div>
                    <div className="link" onClick={() => navigate("/password")}>Change Password</div>
                    <div className="link" onClick={() => onClickLogout()}>Log Out</div>
                </div>
            )}
        </div>
    )
}

export default BurgerMenu;