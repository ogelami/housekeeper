import React from 'react';

class Clock extends React.Component {

  componentDidMount() {
    this.clockInterval = setInterval(this.tick, 50);
    this.clockElement = document.querySelector('.clock > .time');
  }

  componentWillUnmount() {
    clearInterval(this.clockInterval);
  }

  tick = () =>
  {
    this.clockElement.innerHTML = (new Date()).toString();
  };

  render() {
    return (
      <div className="clock"><span className="mdi mdi-clock"/><span className="time"></span></div>
    );
  }
}

export default Clock;
