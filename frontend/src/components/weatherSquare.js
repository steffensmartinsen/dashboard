import {useEffect, useState} from "react";
import { GetWeather, DetermineWeatherIcon } from "../utils/helpers";

const  WeatherSquare = (props) => {
    const {username} = props;
    const [temp, setTemp] = useState(0);
    const [condition, setCondition] = useState("")

    useEffect(() => {
        const fetchWeather = async () => {
            try {
                const data = await GetWeather(username);
                setTemp(data.today.hours[0].temperature)
                setCondition(data.today.hours[0].condition)
            } catch (error) {
                console.error('Error:', error)
            }
        };
        fetchWeather();
    }, [username]);

    return (
        <div className="weather-square">
            <div className="weather-icon">
                {DetermineWeatherIcon(condition)}
            </div>
            <div className="weather-info">
                <div className="weather-temp">{temp}°C</div>
                <div className="weather-desc">{condition}</div>
            </div>
        </div>
    )
}

export default WeatherSquare