import axios from 'axios';
import { NextPageContext } from 'next';
import { getLocale } from 'src/utils/cookie';

const checkContainQuery = (url: string): boolean => {
  const pattern = new RegExp(/\?.+=.*/g);
  return pattern.test(url);
};

export const configureAxios = (
  baseUrl: string,
  ctx?: NextPageContext
): void => {
  axios.defaults.baseURL = baseUrl;
  axios.interceptors.request.use(
    (config) => {
      const newConfig = config;
      const locale = getLocale(ctx);
      const localeQuery = checkContainQuery(newConfig.url)
        ? `&locale=${locale}`
        : `?locale=${locale}`;
      newConfig.url += localeQuery;
      newConfig.headers.Accept = 'application/json';
      newConfig.headers['Content-Type'] = 'application/json';
      return newConfig;
    },
    (error) => {
      return Promise.reject(error);
    }
  );

  axios.interceptors.response.use(
    (response) => response,
    (error) => {
      const { response } = error;
      if (response) {
        console.log(response);
      }
    }
  );
};
