import axios from 'axios';

const TunnelWrap = axiosOptions => {
    let newOptions = Object.assign({}, axiosOptions, {
        url: '//' + (process.env.NODE_ENV === 'development' ? '127.0.0.1' : document.location.host) + '/tun',
        headers: {
            Tunnel: axiosOptions.url
        }
    });

    return axios(newOptions);
};

export default TunnelWrap;

/* TODO: axios tunnel wrapper */

/*
    axios({
      method: 'GET',
      url: 'http://127.0.0.1:80/tun',
      headers: {
        'Tunnel': url
      }
    })
    .then(res => {
      console.log(res);
      this.setState({'disturbanceData': res.data['ResponseData']});
    })
    .catch(function (error) {
      // handle error
      console.log(error);
    });
*/