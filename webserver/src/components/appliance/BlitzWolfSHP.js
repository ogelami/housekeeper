import React from 'react';

class BlitzWolfSHP extends React.Component {
  constructor(props) {
    super(props);

    console.log(props);

    this.state = { on : false };
  }

  config = {
    icon: ['mdi-lightbulb-on-outline', 'mdi-lightbulb-outline'],
    state: 0
  }

  changeState = () => {
    let newState = !this.state.on;

    this.setState({ on: newState });
    this.props.f(this.props.command, newState ? '1' : '0');
  }

  render() {
    return (
      <div onClick={this.changeState}>
        <span className={'mdi ' + (this.state.on ? this.config.icon[0] : this.config.icon[1])}/>
        {this.props.room}, {this.props.location}
      </div>
    );
  }
}

export default BlitzWolfSHP;