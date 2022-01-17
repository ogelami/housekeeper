import ToggleSwitch from '../ToggleSwitch';

export default function BlitzWolfSHP(props) {
  return ToggleSwitch({...{
    icon : ['mdi-lightbulb-on-outline', 'mdi-lightbulb-outline'],
    sendOn : 'ON',
    sendOff : 'OFF',
    receiveOn : 'ON'
  }, ...props});
};
