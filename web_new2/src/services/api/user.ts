import axios from 'axios';
import { generateRequestHeader } from '../utils';
import { IResponse } from 'src/services/response';

interface IUserInfoResponse {
  email: string;
}

class UserService {
  getUserInfo(token: string) {
    return axios.get<null, IResponse<IUserInfoResponse>>('/users/info', {
      headers: generateRequestHeader(token),
    });
  }
}

export default new UserService();
