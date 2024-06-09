import "../App.css";

const Header = () => {

    return (
        <header className="header">
            <nav>
                <ul>
                    <li><a href="/">Home</a></li>
                    <li>|</li>
                    <li><a href="https://github.com/viatheboy/dashboard" target="_blank">Github</a></li>
                    <li>|</li>
                    <li><a href="/">About</a></li>
                </ul>
            </nav>
        </header>
    )
}

export default Header;