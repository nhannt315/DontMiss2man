import * as actionTypes from './actionTypes';

export const login = (email, password, remember) => {
  return {
    type: actionTypes.LOGIN,
    payload: {email, password, remember},
  };
};

export const loginSuccess = (userData, tokenData) => {
  return {
    type: actionTypes.LOGIN_SUCCESS,
    payload: {userData, tokenData},
  };
};

export const loginFailure = error => {
  return {
    type: actionTypes.LOGIN_FAILURE,
    payload: {error},
  };
};

export const register = (email, password, passwordConfirm) => {
  return {
    type: actionTypes.REGISTER,
    payload: {email, password, passwordConfirm},
  };
};

export const registerSuccess = (userData, tokenData) => {
  return {
    type: actionTypes.REGISTER_SUCCESS,
    payload: {userData, tokenData},
  };
};

export const registerFailure = error => {
  return {
    type: actionTypes.REGISTER_FAILURE,
    payload: {error},
  };
};

export const startProcess = () => {
  return {
    type: actionTypes.START_PROCESSING,
  };
};

export const finishProcess = () => {
  return {
    type: actionTypes.FINISH_PROCESSING,
  };
};

export const logout = () => {
  return {
    type: actionTypes.LOGOUT,
  };
};

export const logoutSuccess = () => {
  return {
    type: actionTypes.LOGOUT_SUCCESS,
  };
};

export const logoutFail = () => {
  return {
    type: actionTypes.LOGOUT_FAILURE,
  };
};

export const authCheckState = () => {
  return {
    type: actionTypes.AUTH_CHECK_STATE,
  };
};
