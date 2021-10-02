import axios from 'axios';
import { AccessToken, generateRequestHeader } from '../token';

class AuthService {
  login(email: string, password: string) {
    const payload = { email, password };
    return axios.post('/auth/sign_in', payload);
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
