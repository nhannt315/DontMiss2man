import axios from 'axios';
import {BASE_API_URL} from '../constants/endpoint';

const configureAxios = () => {
  axios.defaults.baseURL = BASE_API_URL;
  axios.interceptors.request.use(config => {
      const newConfig = config;
      newConfig.headers.Accept = 'application/json';
      newConfig.headers['Content-Type'] = 'application/json';
      return newConfig;
    },
    error => {
      return Promise.reject(error);
    },
  );
};

export default configureAxios;
