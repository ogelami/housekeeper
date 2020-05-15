function refreshWeather(chunk)
{
  document.dispatchEvent(new CustomEvent('updateWeather', {'detail' : chunk}));
}

class SimpleSwitch
{
  constructor(selector, command, telemetry) {
    this.element = document.querySelector(selector);
    this.command = command;
    this.telemetry = telemetry;

    this.element.addEventListener('click', () => {
      document.dispatchEvent(new CustomEvent('switchPress', {
        'detail' : {
          'status' : this.element.classList.contains('active'),
          'command' : command
        }
      }));
    });
  }

  telemetryMatch(data) {
    this.element.classList.toggle('active', this.telemetry.off == data);
  }
}

(() => {
  let realTimeClock = document.querySelector('.real-time-clock');

  let weatherIcon = document.querySelector('.weather .icon');
  let weatherDescription = document.querySelector('.weather .description');
  let weatherTemperature = document.querySelector('.weather .temperature');

  let buttonRefreshWeather = document.querySelector('button#refresh-weather');

  let overlay = document.querySelector('div#overlay');
  let overlayOpacitySlider = document.querySelector('input#overlay-opacity-slider');
  let overlayTimeout;

  let consoleLog = document.querySelector('textarea#console-log');
  let consoleLogWebSocket = new WebSocket("ws://" + document.location.host + "/echo");

  let switchList = [
    new SimpleSwitch('#switch-bedroom-left-lamp', {'topic' : 'bedroom_light_left/cmnd/POWER', 'on' : 'ON', 'off' : 'OFF'}, {'topic': 'bedroom_light_left/stat/POWER', 'on' : 'ON', 'off' : 'OFF'}),
    new SimpleSwitch('#switch-kitchen-window-lamp', {'topic' : 'kitchen_light_window/cmnd/POWER', 'on' : 'ON', 'off' : 'OFF'}, {'topic': 'kitchen_light_window/stat/POWER', 'on' : 'ON', 'off' : 'OFF'}),
    new SimpleSwitch('#switch-livingroom-computer-lamp', {'topic' : 'livingroom_light_computer/cmnd/POWER', 'on' : 'ON', 'off' : 'OFF'}, {'topic': 'livingroom_light_computer/stat/POWER', 'on' : 'ON', 'off' : 'OFF'}),
    new SimpleSwitch('#switch-livingroom-window-lamp', {'topic' : 'livingroom_light_window/cmnd/POWER', 'on' : 'ON', 'off' : 'OFF'}, {'topic': 'livingroom_light_window/stat/POWER', 'on' : 'ON', 'off' : 'OFF'}),
    new SimpleSwitch('#switch-livingroom-ceiling-lamp', {'topic' : 'wallace/Formera2510830/switches/0/set', 'on' : '1', 'off' : '0'}, {'topic': 'wallace/Formera2510830/switches/0', 'on' : 'ON', 'off' : 'OFF'}),
    new SimpleSwitch('#switch-kitchen-ceiling-lamp', {'topic' : 'wallace/Formera2508856/switches/1/set', 'on' : '1', 'off' : '0'}, {'topic' : 'wallace/Formera2508856/switches/1', 'on' : '1', 'off' : '0'}),
    new SimpleSwitch('#switch-kitchen-sink-lamp', {'topic' : 'wallace/Formera2508856/switches/0/set', 'on' : '1', 'off' : '0'}, {'topic' : 'wallace/Formera2508856/switches/0', 'on' : '1', 'off' : '0'}),
    new SimpleSwitch('#switch-hallway-ceiling-lamp', {'topic' : 'wallace/Formera2508438/switches/0/set', 'on' : '1', 'off' : '0'}, {'topic' : 'wallace/Formera2508438/switches/0', 'on' : '1', 'off' : '0'}),
    new SimpleSwitch('#switch-bedroom-ceiling-lamp', {'topic' : 'wallace/Formera2508373/switches/0/set', 'on' : '1', 'off' : '0'}, {'topic' : 'wallace/Formera2508373/switches/0', 'on' : '1', 'off' : '0'}),
  ];

  function publishMQTT(topic, message)
  {
    consoleLogWebSocket.send(JSON.stringify({'topic' : topic, 'message' : message}));
  }

  function bumpOverlayTimeout()
  {
    clearTimeout(overlayTimeout);

    overlay.style.display = 'none';

    overlayTimeout = setTimeout(() => {
      overlay.style.display = 'block';
    }, 60000);
  }

  buttonRefreshWeather.addEventListener('click', () => {
    let s = document.createElement('script');
    s.src = 'http://api.openweathermap.org/data/2.5/forecast?lat=59.3399171&lon=17.929966&units=metric&appid=77d3554578ebbd354a066713c397ce84&callback=refreshWeather';
    document.body.appendChild(s);
  });

  overlayOpacitySlider.addEventListener('change', () => {
    overlay.style.opacity = overlayOpacitySlider.value * 0.01;
  });

  //interaction timeout push
  document.addEventListener('touchstart', bumpOverlayTimeout);
  document.addEventListener('mousedown', bumpOverlayTimeout);

  document.addEventListener('updateWeather', (weather) => {
    weatherChunk = weather.detail;

    weatherIcon.src = `/img/openweathermap/${weatherChunk.list[0].weather[0].icon}@2x.png`;
    weatherDescription.innerText = `${weatherChunk.list[0].weather[0].description}`;
    weatherTemperature.innerText = `${weatherChunk.list[0].main.temp} °C`;
  });

  document.addEventListener('switchPress', (state) => 
  {
    console.log(state.detail.command.topic, state.detail.status ? state.detail.command.on : state.detail.command.off);

    publishMQTT(state.detail.command.topic, state.detail.status ? state.detail.command.on : state.detail.command.off);
  });

  consoleLogWebSocket.addEventListener('message', (message) => {
    let messageData = JSON.parse(message.data);

    for (const switchie of switchList)
    {
      if (switchie.telemetry.topic == messageData.topic) {
        switchie.telemetryMatch(messageData.message);
      }
    }
  });

  setInterval(() =>
  {
    realTimeClock.innerText = new Date().toLocaleString("sv-SE", {timeZone: "Europe/Stockholm"});
  }, 500);

  // exposed for debugging
  window.q = consoleLogWebSocket;
  window.pub = publishMQTT;

})();
