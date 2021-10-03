import axios from 'axios';
import { AccessToken, generateRequestHeader } from '../token';
import { IResponse } from 'src/services/response';

interface ILoginRequest {
  email: string;
  password: string;
}

interface ILoginResponseData {
  token: string;
  email: string;
}

class AuthService {
  login(email: string, password: string) {
    const payload = { email, password };
    return axios.post<ILoginRequest, IResponse<ILoginResponseData>>(
      '/auth/login',
      payload
    );
  }

  register(email: string, password: string, passwordConfirm: string) {
    const payload = {
      email,
      password,
      password_confirmation: passwordConfirm,
      confirm_success_url: process.env.REACT_APP_CONFIRM_SUCCESS_URL,
    };

    return axios.post('/auth', payload);
  }

  logout(token: AccessToken) {
    return axios.delete('/auth/sign_out', {
      headers: generateRequestHeader(token),
    });
  }
}

export default new AuthService();
