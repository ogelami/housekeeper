import React from 'react';
import Appliances from './components/Appliances';
import Overlay from './components/Overlay';
import Configuration from './Configuration';

import './App.css';

class App extends React.Component {
  webSocket = new WebSocket('ws://' + (Configuration.webSocketServer || document.location.host) + '/echo');

  broadcastMessage = (topic, message) => {
    if(this.webSocket.readyState === WebSocket.OPEN) {
      this.webSocket.send(JSON.stringify({topic: topic, message: message}));
    }
    else
    {
      console.log('Could not send readyState not open currently set to ' + this.webSocket.readyState);
    }
  }

  componentDidMount() {
    console.log('mounted');

    this.webSocket.onopen = () => {
      console.log('connected');
    }

    this.webSocket.onerror = () => {
      console.log('connected');
    }

    this.webSocket.onclose = () => {
      console.log('erm looks like ws is closing.');
    }
  }

  render() {
    return (
      <div>
        <Overlay/>
        <main>
          <Appliances f={this.broadcastMessage} messageReceived={this.webSocket.onmessage} configuration={Configuration.switchList}/>
        </main>
      </div>
    );
  }
}

export default App;
