import React, { useState } from 'react';
import moment from 'moment';
import tunnelWrap from '../TunnelWrap';
import PropertyValidator, { PropertyValidatorType } from './PropertyValidator';

export default function WeatherBar(props) {
  const validProperties = PropertyValidator(props, {
    /*apiKey : re => re.match(/[a-z0-9]/),*/
    latitude : PropertyValidatorType.float(),
    longitude : PropertyValidatorType.float(),
  });

  const urlParameters = new URLSearchParams({
    'appid': props.apiKey,
    'lat': props.latitude,
    'lon': props.longitude,
    'units': props.units
  });

//  const [weatherData, setWeatherData] = useState([]);
  const [weatherData, setWeatherData] = [[], () => {}];
  console.log(123);

  function fetchWeather() {
    if(!validProperties) {
      return;
    }

//    console.log(`http://api.openweathermap.org/data/2.5/forecast?${urlParameters}`);

    tunnelWrap({url: `http://api.openweathermap.org/data/2.5/forecast?${urlParameters}`}).then(res => {
      setWeatherData(res.data['list'].slice(0, 6).map(r => {
        return {
          'time': moment(r.dt * 1000),
          'icon': `/image/openweather/${r.weather[0].icon}@2x.png`,
          'temperature': Math.round(r.main.temp),
          'humidity': r.main.humidity,
          'wind': Math.round(r.wind.speed)
        };
      }));
    });
  }

  return (
    <div onClick={fetchWeather} className='weather-bar'>
      {weatherData.map((item, iterator) =>
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
};