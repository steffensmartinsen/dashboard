import {useEffect} from "react";
import { GetWeather } from "../utils/helpers";

const  WeatherSquare = (props) => {
    const {username} = props;

    const temp = 20;
    const desc = "Sunny";

    useEffect(() => {
        // Fetch weather data
        GetWeather(username, (data) => {
            console.log(data);
        });
    });

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