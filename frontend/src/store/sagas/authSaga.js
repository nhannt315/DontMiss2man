import {put} from 'redux-saga/effects';
import AuthService from '../../services/authService';
import * as actions from '../actions';
import * as keys from '../../constants/key';

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
    const userData = response.data.data;
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
  const {email, password, passwordConfirm} = action.payload;
  const response = yield AuthService.register(email, password, passwordConfirm);
  yield put(actions.startProcess());
  try {
    const tokenData = {
      accessToken: response.headers['access-token'],
      client: response.headers.client,
      uid: response.headers.uid,
      tokenType: response.headers['token-type'],
      expiry: response.headers.expiry,
    };
    const userData = response.data.data;
    yield localStorage.setItem(keys.TOKEN_DATA_KEY, JSON.stringify(tokenData));
    yield localStorage.setItem(keys.USER_DATA_KEY, JSON.stringify(userData));
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
    yield AuthService.logout(tokenData);
    yield put(actions.logoutSuccess());
  } catch (error) {
    yield put(actions.logoutFail());
  }
}

export function* authCheckStateSaga() {
  const userData = JSON.parse(localStorage.getItem(keys.USER_DATA_KEY));
  const tokenData = JSON.parse(localStorage.getItem(keys.TOKEN_DATA_KEY));
  if (!tokenData) {
    yield put(actions.logoutSuccess());
    return;
  }
  yield put(actions.loginSuccess(userData, tokenData));
}
