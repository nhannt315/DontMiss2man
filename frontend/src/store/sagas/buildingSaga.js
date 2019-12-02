import {put} from 'redux-saga/effects';
import BuildingService from '../../services/buildingService';
import * as actions from '../actions';

export function* fetchBuildingsSaga(action) {
  const {page, perPage, sort} = action.payload;
  yield put(actions.fetchBuildingsStart());
  try {
    const response = yield BuildingService.getBuildingList(page, perPage, sort);
    yield put(actions.fetchBuildingsSuccess(response.list, response.total, response.page));
  } catch (error) {
    yield put(actions.fetchBuildingFailure(error));
  }
}
