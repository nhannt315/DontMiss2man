import {CHANGE_LANGUAGE} from '../actions/actionTypes';

const initialState = {
  language: 'ja',
};

const uiReducer = (state = initialState, {type, payload}) => {
  switch (type) {
    case CHANGE_LANGUAGE:
      return {...state, language: payload.language};
    default:
      return state;
  }
};

export default uiReducer;
