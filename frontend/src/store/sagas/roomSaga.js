import {put} from 'redux-saga/effects';
import RoomService from '../../services/roomService';
import * as actions from '../actions';

export function* fetchRoomDetailSaga(action) {
  const {roomId} = action.payload;
  yield put(actions.fetchRoomStart());
  try {
    const response = yield RoomService.fetchRoomDetail(roomId);
    yield put(actions.fetchRoomSuccess(response));
  } catch (e) {
    yield put(actions.fetchRoomFail(e));
  }
}
