import {put} from 'redux-saga/effects';
import AuthService from '../../services/authService';
import * as actions from '../actions';
import * as keys from '../../constants/key';
import CommonHelper from '../../helpers/common';

export function* loginSaga(action) {
  const {email, password, remember} = action.payload;
  yield put(actions.startProcess());
  try {
    const response = yield AuthService.login(email, password);
    const tokenData = {
      accessToken: response.headers['access-token'],
      client: response.headers.client,
      uid: response.headers.uid,
      tokenType: response.headers['token-type'],
      expiry: response.headers.expiry,
    };
    const userData = response.data;
    if (remember) {
      yield localStorage.setItem(keys.TOKEN_DATA_KEY, JSON.stringify(tokenData));
      yield localStorage.setItem(keys.USER_DATA_KEY, JSON.stringify(userData));
    }
    yield put(actions.loginSuccess(userData, tokenData));
  } catch (error) {
    yield put(actions.loginFailure(error));
  } finally {
    yield put(actions.finishProcess());
  }
}

export function* registerSaga(action) {
  const {email, password, passwordConfirm, callback} = action.payload;
  yield put(actions.startProcess());
  try {
    const response = yield AuthService.register(email, password, passwordConfirm);
    const tokenData = {
      accessToken: response.headers['access-token'],
      client: response.headers.client,
      uid: response.headers.uid,
      tokenType: response.headers['token-type'],
      expiry: response.headers.expiry,
    };
    const userData = response.data;
    if (callback) callback();
    yield put(actions.registerSuccess(userData, tokenData));
  } catch (error) {
    yield put(actions.registerFailure(error));
  } finally {
    yield put(actions.finishProcess());
  }
}

export function* logoutSaga() {
  const tokenData = JSON.parse(localStorage.getItem(keys.TOKEN_DATA_KEY));
  try {
    yield localStorage.removeItem(keys.TOKEN_DATA_KEY);
    yield localStorage.removeItem(keys.USER_DATA_KEY);
    yield put(actions.logoutSuccess());
    yield AuthService.logout(tokenData);
  } catch (error) {
    yield put(actions.logoutFail());
  }
}

export function* authCheckStateSaga() {
  try {
    const userDataStr = localStorage.getItem(keys.USER_DATA_KEY);
    const tokenDataStr = localStorage.getItem(keys.TOKEN_DATA_KEY);
    if (CommonHelper.checkLocalstorageStr(userDataStr) || CommonHelper.checkLocalstorageStr(tokenDataStr)) {
      yield put(actions.logoutSuccess());
      return;
    }
    const userData = JSON.parse(userDataStr);
    const tokenData = JSON.parse(tokenDataStr);
    if (CommonHelper.checkEmptyObject(userData) || CommonHelper.checkEmptyObject(tokenData)) {
      yield put(actions.logoutSuccess());
      return;
    }
    yield put(actions.loginSuccess(userData, tokenData));
  } catch (e) {
    localStorage.removeItem(keys.USER_DATA_KEY);
    localStorage.removeItem(keys.TOKEN_DATA_KEY);
    yield put(actions.logoutSuccess());
  }
}
