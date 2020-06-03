import React from 'react';

class BlitzWolfSHP extends React.Component {
  constructor(props) {
    super(props);

    this.state = { on : false };
  }

  config = {
    icon: ['mdi-lightbulb-on-outline', 'mdi-lightbulb-outline'],
    state: 0
  }

  render() {
    return (
      <div onClick={() => this.setState({on: !this.state.on})}>
        <span className={'mdi ' + (this.state.on ? this.config.icon[0] : this.config.icon[1])}/>
        {this.props.room}, {this.props.location}
      </div>
    );
  }
}

export default BlitzWolfSHP;