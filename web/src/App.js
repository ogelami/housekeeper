import React from 'react';

import Appliances from './components/Appliances';
import Overlay from './components/Overlay';
import WeatherBar from './components/WeatherBar';
import SLDisturbance from './components/SLDisturbance';
import Clock from './components/Clock';

import Configuration from './Configuration';

import './App.css';

class App extends React.Component {
  webSocket = new WebSocket('ws://' + (process.env.NODE_ENV === 'development' ? '127.0.0.1' : document.location.host) + '/echo');
  messageListeners = [];

  broadcastMessage = (topic, message) => {
    if(this.webSocket.readyState === WebSocket.OPEN) {
      this.webSocket.send(JSON.stringify({topic: topic, message: message}));
    }
    else {
      console.log('Could not send readyState not open currently set to ' + this.webSocket.readyState);
    }
  }

  registerMessageReceivedListener = func => {
    this.messageListeners.push(func);
  }

  componentDidMount() {
    this.webSocket.onopen = (q) => {
      console.log('connected');
    }

    this.webSocket.onerror = () => {
      console.log('Connecting to the websocket server has failed, is the backend running?');
    }

    this.webSocket.onclose = () => {
      console.log('Websocket is now closing.');
    }

    this.webSocket.onmessage = (message) => {
      for (const listener of this.messageListeners) {
        listener.call(this, JSON.parse(message.data));
      }
    }
  }

  render() {
    return (
      <div>
        <Overlay/>
        <main>
          <nav><Clock /></nav>
          <div className="appliances">
            <Appliances broadcastMessage={this.broadcastMessage} registerMessageReceivedListener={this.registerMessageReceivedListener} configuration={Configuration.switchList}/>
          </div>
          <WeatherBar apiKey={Configuration.openWeather.apiKey} longitude={Configuration.openWeather.longitude} latitude={Configuration.openWeather.latitude} refreshRate={Configuration.openWeather.refreshRate} units={Configuration.openWeather.units}/>
          <SLDisturbance apiKey={Configuration.sLDisturbance.apiKey} transportMode={Configuration.sLDisturbance.transportMode} lineNumber={Configuration.sLDisturbance.lineNumber} refreshRate={Configuration.sLDisturbance.refreshRate} />
        </main>
      </div>
    );
  }
}

export default App;
