import React from 'react';
import Appliances from './components/Appliances';
import Overlay from './components/Overlay';
import Configuration from './Configuration';

import './App.css';

class App extends React.Component {
  webSocket = new WebSocket('ws://' + (Configuration.webSocketServer || document.location.host) + '/echo');
  messageListeners = [];

  broadcastMessage = (topic, message) => {
    if(this.webSocket.readyState === WebSocket.OPEN) {
      this.webSocket.send(JSON.stringify({topic: topic, message: message}));
    }
    else
    {
      console.log('Could not send readyState not open currently set to ' + this.webSocket.readyState);
    }
  }

  registerMessageReceivedListener = func => {
    this.messageListeners.push(func);
  }

  componentDidMount() {
    this.webSocket.onopen = () => {
      console.log('connected');
    }

    this.webSocket.onerror = () => {
      console.log('connected');
    }

    this.webSocket.onclose = () => {
      console.log('erm looks like ws is closing.');
    }

    this.webSocket.onmessage = (message) => {
      for (const listener of this.messageListeners) {
        listener.call(this, JSON.parse(message.data));
//        console.log(listener);
      }
    }
  }

  render() {
    return (
      <div>
        <Overlay/>
        <main>
          <div className="appliances">
            <Appliances broadcastMessage={this.broadcastMessage} registerMessageReceivedListener={this.registerMessageReceivedListener} configuration={Configuration.switchList}/>
          </div>
        </main>
      </div>
    );
  }
}

export default App;
