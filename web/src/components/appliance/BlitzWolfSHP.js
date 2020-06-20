import React from 'react';

class BlitzWolfSHP extends React.Component {
  constructor(props) {
    super(props);
    this.state = { on : false };
  }

  componentDidMount() {
    this.props.registerMessageReceivedListener(data => {
      if (data.topic === this.props.status) {
        this.setState({ on: data.message === 'ON' });
        console.log(data.message);
      }
    });
  }

  config = {
    icon: ['mdi-lightbulb-on-outline', 'mdi-lightbulb-outline'],
    state: 0
  }

  changeState = () => {
    let newState = !this.state.on;

    this.props.broadcastMessage(this.props.command, newState ? 'ON' : 'OFF');
  }

  render() {
    return (
      <div onClick={this.changeState} className={'flip-card ' + (this.state.on ? 'on':'')}>
        <div className="flip-card-inner">
          <div className="flip-card-front">
            <span className={'mdi ' + this.config.icon[1]}/>
            {this.props.room}, {this.props.location}
          </div>
          <div className="flip-card-back">
            <span className={'mdi ' + this.config.icon[0]}/>
            {this.props.room}, {this.props.location}
          </div>
        </div>
      </div>
    );
  }
}

export default BlitzWolfSHP;
