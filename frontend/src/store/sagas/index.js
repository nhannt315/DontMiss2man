import {all, takeLatest} from 'redux-saga/effects';
import * as actionTypes from '../actions/actionTypes';
import {fetchBuildingsSaga} from './buildingSaga';
import {fetchRoomDetailSaga} from './roomSaga';
import {loginSaga, registerSaga, logoutSaga, authCheckStateSaga} from './authSaga';

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

export function* watchAuth() {
  yield all([
    takeLatest(actionTypes.LOGIN, loginSaga),
    takeLatest(actionTypes.REGISTER, registerSaga),
    takeLatest(actionTypes.LOGOUT, logoutSaga),
    takeLatest(actionTypes.AUTH_CHECK_STATE, authCheckStateSaga),
  ]);
}
