import axios from 'axios';

const AuthService = {};

const getHeader = tokenData => ({
  'access-token': tokenData.accessToken,
  'token-type': tokenData.tokenType,
  uid: tokenData.uid,
  client: tokenData.client,
});

AuthService.login = (email, password) => {
  const payload = {email, password};
  return axios.post('/auth/sign_in', payload);
};

AuthService.register = (email, password, passwordConfirm) => {
  const payload = {
    email, password,
    password_confirmation: passwordConfirm,
    confirm_success_url: process.env.REACT_APP_CONFIRM_SUCCESS_URL,
  };
  return axios.post('/auth', payload);
};

AuthService.logout = tokenData => axios.delete('/auth/sign_out', {headers: getHeader(tokenData)});

export default AuthService;
