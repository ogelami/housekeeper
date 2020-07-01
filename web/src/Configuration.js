import BlitzWolfSHP from './components/appliance/BlitzWolfSHP';
import FormeraWall from './components/appliance/FormeraWall';

const Config = {
  webSocketServer: '127.0.0.1',
  sLDisturbance:
  {
    apiKey: '',
    transportMode: 'metro',
    lineNumber: [17, 18, 19],
    refreshRate: 600000 /* milliseconds or 0 for on request  */
  },
  openWeather:
  {
    apiKey: '',
    latitude: 0.0,
    longitude: 0.0,
    refreshRate: 600000,
    units: 'metric'
  },
  switchList: [
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
      location: 1,
      command: 'livingroom_light_window/cmnd/POWER',
      status: 'livingroom_light_window/stat/POWER'
    },
    {
      type: BlitzWolfSHP,
      room: 'bedroom',
      location: 'left',
      command: 'bedroom_light_left/cmnd/POWER',
      status: 'bedroom_light_left/cmnd/POWER'
    },
    {
      type: BlitzWolfSHP,
      room: 'bedroom',
      location: 'window',
      command: 'bedroom_light_window/cmnd/POWER',
      status: 'bedroom_light_window/cmnd/POWER'
    },
    {
      type: FormeraWall,
      room: 'livingroom',
      location: 'window',
      command: 'livingroom_light_window/cmnd/POWER',
      status: 'livingroom_light_window/stat/POWER'
    }
  ]
};

export default Config;
