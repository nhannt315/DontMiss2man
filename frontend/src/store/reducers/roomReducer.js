import * as actionTypes from '../actions/actionTypes';

const initialState = {
  room: {},
  loading: true,
  error: null,
};

const roomReducer = (state = initialState, {type, payload}) => {
  switch (type) {
    case actionTypes.FETCH_ROOM_START:
      return {...state, loading: true};
    case actionTypes.FETCH_ROOM_SUCCESS:
      return {...state, room: payload.room, loading: false};
    case actionTypes.FETCH_ROOM_FAILURE:
      return {...state, error: payload.error, loading: false};
    default:
      return state;
  }
};

export default roomReducer;
