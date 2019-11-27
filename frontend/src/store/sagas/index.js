import {all, takeLatest} from 'redux-saga/effects';
import * as actionTypes from '../actions/actionTypes';
import {fetchBuildingsSaga} from './buildingSaga';

export function* watchBuilding() {
  yield all([
    takeLatest(actionTypes.FETCH_BUILDINGS, fetchBuildingsSaga),
  ]);
}
