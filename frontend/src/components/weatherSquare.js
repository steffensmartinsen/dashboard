import {useEffect, useState} from "react";
import { GetWeather, DetermineWeatherIcon } from "../utils/helpers";

const  WeatherSquare = (props) => {
    const {username} = props;
    const [temp, setTemp] = useState(0);
    const [condition, setCondition] = useState("")
    const [city, setCity] = useState("")

    useEffect(() => {
        const fetchWeather = async () => {
            try {
                const data = await GetWeather(username);
                setTemp(data.today.hours[0].temperature)
                setCondition(data.today.hours[0].condition)
                setCity(data.city)
            } catch (error) {
                console.error('Error:', error)
            }
        };
        fetchWeather();
    }, [username]);

    return (
        <div className="weatherContainer">
            <div className="weather-square">
                <div className="weather-city">{city}</div>
                <div className="weather-icon">
                    {DetermineWeatherIcon(condition)}
                </div>
                <div className="weather-info">
                    <div className="weather-temp">{temp}°C</div>
                    <div className="weather-desc">{condition}</div>
                </div>
            </div>
        </div>
    )
}

export default WeatherSquare