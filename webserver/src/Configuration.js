import BlitzWolfSHP from './components/appliance/BlitzWolfSHP';

const SwitchList = [
  {
    type: BlitzWolfSHP,
    room: 'livingroom',
    location: 'computer',
    command: 'livingroom_light_computer/cmnd/POWER',
    status: 'livingroom_light_computer/stat/POWER'
  },
  {
    type: BlitzWolfSHP,
    room: 'livingroom',
    location: 'window',
    command: 'livingroom_light_window/cmnd/POWER',
    status: 'livingroom_light_window/stat/POWER'
  }
];

export default SwitchList;