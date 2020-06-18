import React from 'react';

class Applicances extends React.Component {
  render() {
    return this.props.configuration.map((_switch, iterator) => {
      const SwitchType = _switch.type;

      return <SwitchType broadcastMessage={this.props.broadcastMessage} registerMessageReceivedListener={this.props.registerMessageReceivedListener} key={iterator} command={_switch.command} status={_switch.status} room={_switch.room} location={_switch.location} />
    });
  }
}

export default Applicances;
