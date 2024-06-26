
function WeatherSquare() {

    const temp = 20;
    const desc = "Sunny";


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