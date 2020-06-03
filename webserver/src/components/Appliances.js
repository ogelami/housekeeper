import React from 'react';

class Applicances extends React.Component {
  render() {
    return this.props.configuration.map((_switch, iterator) => {
      const SwitchType = _switch.type;

      return <SwitchType key={iterator} room={_switch.room} location={_switch.location} />
    });
  }
}

export default Applicances;
