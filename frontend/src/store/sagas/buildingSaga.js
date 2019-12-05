import {put} from 'redux-saga/effects';
import BuildingService from '../../services/buildingService';
import * as actions from '../actions';

export function* fetchBuildingsSaga(action) {
  const {page, perPage, sort, condition} = action.payload;
  yield put(actions.fetchBuildingsStart(sort, condition, perPage));
  try {
    const response = yield BuildingService.getBuildingList(page, perPage, sort, condition);
    yield put(actions.fetchBuildingsSuccess(response.list, response.total, response.page));
  } catch (error) {
    yield put(actions.fetchBuildingFailure(error));
  }
}
