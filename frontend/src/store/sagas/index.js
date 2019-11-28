import {all, takeLatest} from 'redux-saga/effects';
import * as actionTypes from '../actions/actionTypes';
import {fetchBuildingsSaga} from './buildingSaga';
import {fetchRoomDetailSaga} from './roomSaga';

export function* watchBuilding() {
  yield all([
    takeLatest(actionTypes.FETCH_BUILDINGS_REQUEST, fetchBuildingsSaga),
  ]);
}

export function* watchRoom() {
  yield all([
    takeLatest(actionTypes.FETCH_ROOM_REQUEST, fetchRoomDetailSaga),
  ]);
}
