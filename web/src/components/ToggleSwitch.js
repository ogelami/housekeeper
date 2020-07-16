import React, { useState } from 'react';
import PropertyValidator, { PropertyValidatorType } from './PropertyValidator';

export default function ToggleSwitch(props) {
  PropertyValidator(props, {
    room : PropertyValidatorType.string(),
    location : PropertyValidatorType.string(),
    command : PropertyValidatorType.string(),
    status : PropertyValidatorType.string()
  });

  const [currentState, setState] = useState(false);

  const toggleSwitch = () => props.broadcastMessage(props.command, !currentState ? props.sendOn : props.sendOff);

  props.registerMessageReceivedListener(data => {
      if(data.topic === props.status) {
        setState(data.message === props.receiveOn);
      }
  });

  props = {...{
    icon : ['mdi-lightbulb-on-outline', 'mdi-lightbulb-outline'],
    sendOn : 'ON',
    sendOff : 'OFF',
    receiveOn : 'ON'
  }, ...props};

  return (
    <div onClick={toggleSwitch} className={'flip-card ' + (currentState ? 'on':'')}>
      <div className="flip-card-inner">
        <div className="flip-card-front">
          <span className={'mdi ' + props.icon[1]}/>
          {props.room}<br/> {props.location}
        </div>
        <div className="flip-card-back">
          <span className={'mdi ' + props.icon[0]}/>
          {props.room}<br/> {props.location}
        </div>
      </div>
    </div>
  );
};
