import React from 'react';
import Configuration from '../Configuration'

export default function Applicances(props) {
  return(
    <div className="appliances">
      {
        Configuration.switchList.map((_switch, iterator) => {
          const SwitchType = _switch.type;

          return <SwitchType broadcastMessage={props.broadcastMessage} registerMessageReceivedListener={props.registerMessageReceivedListener} key={iterator} command={_switch.command} status={_switch.status} room={_switch.room} location={_switch.location} />
        })
      }
    </div>
  );
};