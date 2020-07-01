import React from 'react';
import SuperComponent from './SuperComponent';

class ToggleSwitch extends SuperComponent {
    constructor(props) {
      super(props,{
        room : SuperComponent.availablePropTypes.string,
        location : SuperComponent.availablePropTypes.string,
        command : SuperComponent.availablePropTypes.string,
        status : SuperComponent.availablePropTypes.string
      });
      
      this.validatePropTypes();

      this.state = { on : false };
    }
  
    componentDidMount() {
        this.props.registerMessageReceivedListener(data => {
            if(data.topic === this.props.status) {
              this.setState({ on : data.message === this.receiveOn});
            }
        });
    }

    toggle = () => {
      let newState = !this.state.on;

      this.props.broadcastMessage(this.props.command, newState ? this.sendOn : this.sendOff);
    }
    
    render() {
        return (
          <div onClick={this.toggle} className={'flip-card ' + (this.state.on ? 'on':'')}>
            <div className="flip-card-inner">
              <div className="flip-card-front">
                <span className={'mdi ' + this.icon[1]}/>
                {this.props.room}, {this.props.location}
              </div>
              <div className="flip-card-back">
                <span className={'mdi ' + this.icon[0]}/>
                {this.props.room}, {this.props.location}
              </div>
            </div>
          </div>
        );
      }
}

export default ToggleSwitch;
