import {useEffect} from "react";

const  WeatherSquare = (props) => {
    const {username} = props;

    const temp = 20;
    const desc = "Sunny";

    useEffect(() => {
        // Fetch weather data
        fetch("http://localhost:8080/dashboards/v1/weather/" + username + "/")

    }, []);

    return (
        <div className="weather-square">
            <div className="weather-icon">
                Placeholder Weather Icon
                <img alt="weather icon" />
            </div>
            <div className="weather-info">
                <div className="weather-temp">{temp}°C</div>
                <div className="weather-desc">{desc}</div>
            </div>
        </div>
    )
}

export default WeatherSquare