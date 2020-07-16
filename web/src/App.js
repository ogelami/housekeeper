import React from 'react';

import Appliances from './components/Appliances';
import Overlay from './components/Overlay';
import WeatherBar from './components/WeatherBar';
import SLDisturbance from './components/SLDisturbance';
import Clock from './components/Clock';

import Configuration from './Configuration';

import './App.css';

export default function App() {
  const webSocket = new WebSocket('ws://' + (process.env.NODE_ENV === 'development' ? '127.0.0.1' : document.location.host) + '/echo');
  const messageListeners = [];

  function broadcastMessage(topic, message) {
    if(webSocket.readyState === WebSocket.OPEN) {
      webSocket.send(JSON.stringify({topic: topic, message: message}));
    } else {
      console.log('Could not send readyState not open currently set to ' + webSocket.readyState);
    }
  }

  function registerMessageReceivedListener(func) {
    messageListeners.push(func);
  }

  webSocket.onopen = (q) => {
    console.log('connected');
  }

  webSocket.onerror = () => {
    console.log('Connecting to the websocket server has failed, is the backend running?');
  }

  webSocket.onclose = () => {
    console.log('Websocket is now closing.');
  }

  webSocket.onmessage = (message) => {
    for (const listener of messageListeners) {
      listener.call(this, JSON.parse(message.data));
    }
  }

  return (
    <div>
      <Overlay/>
      <main>
        <nav><Clock /></nav>
        <Appliances broadcastMessage={broadcastMessage} registerMessageReceivedListener={registerMessageReceivedListener}/>
        <WeatherBar apiKey={Configuration.openWeather.apiKey} longitude={Configuration.openWeather.longitude} latitude={Configuration.openWeather.latitude} refreshRate={Configuration.openWeather.refreshRate} units={Configuration.openWeather.units}/>
        <SLDisturbance apiKey={Configuration.sLDisturbance.apiKey} transportMode={Configuration.sLDisturbance.transportMode} lineNumber={Configuration.sLDisturbance.lineNumber} refreshRate={Configuration.sLDisturbance.refreshRate} />
      </main>
    </div>
  );
};
