import React from 'react';
import axios from 'axios';
import moment from 'moment';
import SuperComponent from './SuperComponent';

class WeatherBar extends SuperComponent {
  constructor(props) {
    super(props, {
      apiKey : re => re.match(/[a-zA-Z0-9]/),
      latitude : SuperComponent.availablePropTypes.float,
      longitude : SuperComponent.availablePropTypes.float
    });

    this.state = { weatherData : [] };
  }
  
  componentDidMount() {
    this.refreshInterval = setInterval(this.fetchWeather, this.props.refreshRate);
    this.fetchWeather();
  }

  componentWillUnmount() {
    clearInterval(this.refreshInterval);
  }

  fetchWeather = () => {
    if(!this.validatePropTypes())
    {
      return;
    }

    const urlParameters = new URLSearchParams({
      'appid': this.props.apiKey,
      'lat': this.props.latitude,
      'lon': this.props.longitude,
      'units': this.props.units
    });

//    console.log(`http://api.openweathermap.org/data/2.5/forecast?${urlParameters}`);

    this.setState({'weatherData':[]});

    axios.get(`http://api.openweathermap.org/data/2.5/forecast?${urlParameters}`)
      .then(res => {
        const weatherData = res.data;

        this.setState({'weatherData': weatherData['list'].slice(0, 4).map(r => {
          return {
            'time': moment(r.dt * 1000),
            'icon': `/image/openweather/${r.weather[0].icon}@2x.png`,
            'temperature': Math.round(r.main.temp),
            'humidity': r.main.humidity,
            'wind': Math.round(r.wind.speed)
          };
        })});
      });
  }

  render() {
    return (
      <div onClick={this.fetchWeather} className='weather-bar'>
        {this.state.weatherData.map((item, iterator) =>
          <div className='weather-block' key={iterator}>
            <span className="time">{item.time.format('HH:mm')}</span>
            <img alt={'image for weather condition ' + iterator} src={item.icon} />
            <span className="temperature">{item.temperature} °C</span>
            <span className="humidity">{item.humidity} %</span>
            <span className="wind">{item.wind} ms</span>
          </div>
        )}
      </div>
    );
  }
}

export default WeatherBar;
