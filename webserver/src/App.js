import React from 'react';
import Appliances from './components/Appliances';
import Overlay from './components/Overlay';
import Configuration from './Configuration';

import './App.css';

class App extends React.Component {
  webSocket = new WebSocket('ws://' + document.location.host + '/echo');

  func = state => {
    this.webSocket.send('oh, hi mark' + (state ? '1':'0'));
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
          <Appliances f={this.func} configuration={Configuration}/>
        </main>
      </div>
    );
  }
}

export default App;
