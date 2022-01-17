import ToggleSwitch from '../ToggleSwitch';

export default function FormeraWall(props) {
  return ToggleSwitch({...{
    icon : ['mdi-lightbulb-on-outline', 'mdi-lightbulb-outline'],
    sendOn : '1',
    sendOff : '0',
    receiveOn : '1'
  }, ...props});
};
