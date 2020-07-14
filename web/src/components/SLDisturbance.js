//https://api.sl.se/api2/deviations.JSON?key=750131196e424fadb988372cb83df24c&transportMode=metro&lineNumber=17,18,19

import React from 'react';
import tunnelWrap from '../TunnelWrap';
import SuperComponent from './SuperComponent';

class SLDisturbance extends SuperComponent {
  constructor(props) {
    super(props, { apiKey : re => re.match(/[a-z0-9]/) });
    this.state = { disturbanceData : [] };
  }

  componentDidMount() {
    if(this.props.refreshRate) {
      this.refreshInterval = setInterval(this.fetchSLDisturbances, this.props.refreshRate);
    }

    this.fetchSLDisturbances();
  }

  componentWillUnmount() {
    clearInterval(this.refreshInterval);
  }

  fetchSLDisturbances = () => {
    if(!this.validatePropTypes()) {
      return;
    }

    const urlParameters = new URLSearchParams({ 'key': this.props.apiKey });

    if(this.props.transportMode) {
      urlParameters.append('transportMode', this.props.transportMode);
    }

    if(this.props.lineNumber) {
      urlParameters.append('lineNumber', this.props.lineNumber.join(','));
    }

    const url = `http://api.sl.se/api2/deviations.JSON?${urlParameters}`;

    this.setState({'disturbanceData': []});

    tunnelWrap({
      method: 'GET',
      url: url
    })
    .then(res => {
      let disturbanceData = res.data['ResponseData'].map(r => Object.assign((({ Scope, Header, Details }) => ({ Scope, Header, Details }))(r),{open: false}))

      this.setState({'disturbanceData': disturbanceData});
    })
    .catch(function (error) {
      // handle error
      console.log(error);
    });
  }

  toggleDisplay = (index) => {
/*    console.log('displaying', index);
*/
    document.querySelectorAll(`div.sl-disturbance div div.details`)[index].classList.toggle('open');
  }

  render() {
    return (
      <div /*onClick={this.fetchSLDisturbances}*/ className='sl-disturbance'>
        {this.state.disturbanceData.map((item, index) => 
          <div key={index}>
            <div className='header' onClick={() => this.toggleDisplay(index)}>
              <h4>{item['Scope']} - {item['Header']}</h4>
            </div>
            <div className={'details ' + (item['open'] ? 'open':'')}>
              <p>{item['Details']}</p>
            </div>
          </div>
        )}
      </div>
    );
  }
}

export default SLDisturbance;