import axios from 'axios';

const RoomService = {};

RoomService.fetchRoomDetail = roomId => {
  return axios.get(`/rooms/${roomId}`).then(res => res.data.data);
};

export default RoomService;
