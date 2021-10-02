import * as actionTypes from './actionTypes';

export const fetchBuildings = (page, perPage, sort, condition) => {
  return {
    type: actionTypes.FETCH_BUILDINGS_REQUEST,
    payload: {page, perPage, sort, condition},
  };
};

export const fetchBuildingsStart = (sort, condition, perPage) => {
  return {
    type: actionTypes.FETCH_BUILDINGS_START,
    payload: {sort, condition, perPage},
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
