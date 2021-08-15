import {
  FETCH_BUILDINGS_START,
  FETCH_BUILDINGS_FAILURE,
  FETCH_BUILDINGS_SUCCESS,
} from '../actions/actionTypes';

const initialState = {
  list: [],
  totalCount: 1,
  currentPage: 1,
  error: null,
  loading: false,
  isUseSaveState: false,
  sort: null,
  condition: null,
  perPage: null,
};

const buildingReducer = (state = initialState, {type, payload}) => {
  switch (type) {
    case FETCH_BUILDINGS_START:
      return {...state, loading: true, sort: payload.sort, condition: payload.condition, perPage: payload.perPage};
    case FETCH_BUILDINGS_SUCCESS:
      return {...state, list: payload.list, currentPage: parseInt(payload.page, 10), totalCount: payload.totalCount, loading: false};
    case FETCH_BUILDINGS_FAILURE:
      return {...state, error: payload.error, loading: false};
    default:
      return state;
  }
};

export default buildingReducer;
