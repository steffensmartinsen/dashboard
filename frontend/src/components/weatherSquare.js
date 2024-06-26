
const weatherSquare = () => {

    return (
        <div className="weather-square">
            <div className="weather-icon">
                <img src={icon} alt="weather icon" />
            </div>
            <div className="weather-info">
                <div className="weather-temp">{temp}°C</div>
                <div className="weather-desc">{desc}</div>
            </div>
        </div>
    )
}

export default weatherSquare