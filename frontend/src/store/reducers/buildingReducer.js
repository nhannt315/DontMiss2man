import {
  FETCH_BUILDINGS_START,
  FETCH_BUILDINGS_FAILURE,
  FETCH_BUILDINGS_SUCCESS,
} from '../actions/actionTypes';

const initialState = {
  list: [],
  totalPages: 1,
  currentPage: 1,
  error: null,
  loading: false,
};

const buildingReducer = (state = initialState, {type, payload}) => {
  switch (type) {
    case FETCH_BUILDINGS_START:
      return {...state, loading: true};
    case FETCH_BUILDINGS_SUCCESS:
      return {...state, list: payload.list, currentPage: payload.page, totalPages: payload.totalPages, loading: false};
    case FETCH_BUILDINGS_FAILURE:
      return {...state, error: payload.error, loading: false};
    default:
      return state;
  }
};

export default buildingReducer;
