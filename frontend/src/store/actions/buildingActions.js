import * as actionTypes from './actionTypes';

export const fetchBuildings = (page, perPage) => {
  return {
    type: actionTypes.FETCH_BUILDINGS_REQUEST,
    payload: {page, perPage},
  };
};

export const fetchBuildingsStart = () => {
  return {
    type: actionTypes.FETCH_BUILDINGS_START,
  };
};

export const fetchBuildingsSuccess = (list, totalCount, page) => {
  return {
    type: actionTypes.FETCH_BUILDINGS_SUCCESS,
    payload: {list, totalCount, page},
  };
};

export const fetchBuildingFailure = error => {
  return {
    type: actionTypes.FETCH_BUILDINGS_FAILURE,
    payload: {error},
  };
};
