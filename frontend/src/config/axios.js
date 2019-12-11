import axios from 'axios';
import {message} from 'antd';
import get from 'lodash/get';
import {BASE_API_URL} from '../constants/endpoint';
import {LOCALE_KEY} from '../constants/key';

const checkContainQuery = url => {
  const pattern = new RegExp(/\?.+=.*/g);
  return pattern.test(url);
};

const configureAxios = () => {
  axios.defaults.baseURL = BASE_API_URL;
  axios.interceptors.request.use(config => {
      const newConfig = config;
      const locale = localStorage.getItem(LOCALE_KEY) || 'ja';
      const localeQuery = checkContainQuery(newConfig.url) ? `&locale=${locale}` : `?locale=${locale}`;
      newConfig.url += localeQuery;
      newConfig.headers.Accept = 'application/json';
      newConfig.headers['Content-Type'] = 'application/json';
      return newConfig;
    },
    error => {
      return Promise.reject(error);
    },
  );


  axios.interceptors.response.use(response => {
    const newRes = response;
    const resData = response.data;
    newRes.data = resData.data;
    return newRes;
  }, error => {
    const {response} = error;
    if (response) {
      switch (response.status) {
        case 401:
          message.error(get(response, 'data.errors[0]'));
          break;
        case 422:
          const errors = get(response, 'data.errors');
          const errorMessages = Object.keys(errors).map(key => {
            if (key !== 'full_messages')
              return errors[key];
            return '';
          });
          message.error(errorMessages.join(''));
          break;
        default:
          break;
      }
    }
  });
};

export default configureAxios;
