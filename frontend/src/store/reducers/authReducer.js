import {
  LOGIN_SUCCESS,
  LOGIN_FAILURE,
  REGISTER_SUCCESS,
  REGISTER_FAILURE,
  LOGOUT_SUCCESS,
  LOGOUT_FAILURE,
  START_PROCESSING,
  FINISH_PROCESSING,
} from '../actions/actionTypes';

const initialState = {
  isAuthenticated: false,
  userData: {},
  tokenData: {},
  error: null,
  fbUserData: null,
  credential: null,
  isProcessing: false,
};

const authReducer = (state = initialState, {type, payload}) => {
  switch (type) {
    case START_PROCESSING:
      return {...state, isProcessing: true};
    case FINISH_PROCESSING:
      return {...state, isProcessing: false};
    case REGISTER_SUCCESS:
    case LOGIN_SUCCESS:
      return {...state, isAuthenticated: true, userData: payload.userData, tokenData: payload.tokenData};
    case REGISTER_FAILURE:
    case LOGIN_FAILURE:
      return {...state, isAuthenticated: false, error: payload.error};
    case LOGOUT_SUCCESS:
    case LOGOUT_FAILURE:
      return initialState;
    default:
      return state;
  }
};

export default authReducer;
