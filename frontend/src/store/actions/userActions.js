import {ADD_USER_FAVORITE, REMOVE_USER_FAVORITE} from './actionTypes';

export const addUserFavorite = roomId => {
  return {
    type: ADD_USER_FAVORITE,
    payload: {roomId},
  };
};

export const removeUserFavorite = roomId => {
  return {
    type: REMOVE_USER_FAVORITE,
    payload: {roomId},
  };
};
