import {
  LOGIN_SUCCESS,
  LOGIN_FAILURE,
  REGISTER_SUCCESS,
  REGISTER_FAILURE,
  LOGOUT_SUCCESS,
  LOGOUT_FAILURE,
  START_PROCESSING,
  FINISH_PROCESSING,
  ADD_USER_FAVORITE,
  REMOVE_USER_FAVORITE,
} from '../actions/actionTypes';

import * as keys from '../../constants/key';

const initialState = {
  isAuthenticated: false,
  userData: {},
  tokenData: {},
  error: null,
  isProcessing: false,
};

const addUserFavorite = (state, payload) => {
  const newList = state.userData.favorites;
  newList.push(payload.roomId);
  const newUserData = {...state.userData, favorites: newList};
  localStorage.setItem(keys.USER_DATA_KEY, JSON.stringify(newUserData));
  return {...state, userData: newUserData};
};

const removeUserFavorite = (state, payload) => {
  const newList = state.userData.favorites.filter(e => e !== payload.roomId);
  const newUserData = {...state.userData, favorites: newList};
  localStorage.setItem(keys.USER_DATA_KEY, JSON.stringify(newUserData));
  return {...state, userData: newUserData};
};

const authReducer = (state = initialState, {type, payload}) => {
  switch (type) {
    case START_PROCESSING:
      return {...state, isProcessing: true};
    case FINISH_PROCESSING:
      return {...state, isProcessing: false};
    case REGISTER_SUCCESS:
      return state;
    case LOGIN_SUCCESS:
      return {...state, isAuthenticated: true, userData: payload.userData, tokenData: payload.tokenData};
    case REGISTER_FAILURE:
    case LOGIN_FAILURE:
      return {...state, isAuthenticated: false, error: payload.error};
    case LOGOUT_SUCCESS:
    case LOGOUT_FAILURE:
      localStorage.removeItem(keys.USER_DATA_KEY);
      localStorage.removeItem(keys.TOKEN_DATA_KEY);
      return initialState;
    case ADD_USER_FAVORITE:
      return addUserFavorite(state, payload);
    case REMOVE_USER_FAVORITE:
      return removeUserFavorite(state, payload);
    default:
      return state;
  }
};


export default authReducer;
