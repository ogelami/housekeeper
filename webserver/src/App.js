import React from 'react';
import Appliances from './components/Appliances';
import Overlay from './components/Overlay';
import Configuration from './Configuration';

import './App.css';

class App extends React.Component {

  consoleLogWebSocket = new WebSocket('ws://' + document.location.host + '/echo');

  render() {
    return (
      <div>
        <Overlay/>
        <main>
          <Appliances configuration={Configuration}/>
        </main>
      </div>
    );
  }
}

export default App;
