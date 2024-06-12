import { useState } from "react";
import { GiHamburgerMenu } from "react-icons/gi";
import { IoMdClose } from "react-icons/io";
import { Logout } from "../utils/helpers";

const BurgerMenu = (props) => {
    const [menuOpen, setMenuOpen] = useState(false);

    const handleMenuClick = () => {
        setMenuOpen(!menuOpen);
    }

    return (
        <div className="burgerMenu">
            <div className="burgerMenuIcon" onClick={handleMenuClick}>
                <GiHamburgerMenu />
            </div>
            {menuOpen && (
                <div className="burgerMenuContent">
                    <div className="closeIcon" onClick={handleMenuClick}>
                        <IoMdClose />
                    </div>
                    <a href="/account-settings">Edit Account</a>
                    <a href="/preferences">Edit Preferences</a>
                    <a href="#" onClick={() => Logout(props.username, props.setLoggedIn)}>Log Out</a>
                </div>
            )}
        </div>
    )
}

export default BurgerMenu;