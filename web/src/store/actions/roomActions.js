import * as actionTypes from './actionTypes';

export const fetchRoom = id => {
  return {
    type: actionTypes.FETCH_ROOM_REQUEST,
    payload: {
      roomId: id,
    },
  };
};

export const fetchRoomStart = () => {
  return {
    type: actionTypes.FETCH_ROOM_START,
  };
};

export const fetchRoomSuccess = room => {
  return {
    type: actionTypes.FETCH_ROOM_SUCCESS,
    payload: {room},
  };
};

export const fetchRoomFail = error => {
  return {
    type: actionTypes.FETCH_ROOM_FAILURE,
    payload: {error,},
  };
};
