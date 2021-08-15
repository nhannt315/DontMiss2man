import axios from 'axios';

const RoomService = {};

RoomService.fetchRoomDetail = roomId => {
  return axios.get(`/rooms/${roomId}`);
};

export default RoomService;
