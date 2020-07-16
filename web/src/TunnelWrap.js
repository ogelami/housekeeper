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
