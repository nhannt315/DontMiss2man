import * as actionTypes from './actionTypes';

export const fetchBuildings = (page, perPage) => {
  return {
    type: actionTypes.FETCH_BUILDINGS,
    payload: {page, perPage},
  };
};

export const fetchBuildingsStart = () => {
  return {
    type: actionTypes.FETCH_BUILDINGS_START,
  };
};

export const fetchBuildingsSuccess = (list, totalPages, page) => {
  return {
    type: actionTypes.FETCH_BUILDINGS_SUCCESS,
    payload: {list, totalPages, page},
  };
};

export const fetchBuildingFailure = error => {
  return {
    type: actionTypes.FETCH_BUILDINGS_FAILURE,
    payload: {error},
  };
};
