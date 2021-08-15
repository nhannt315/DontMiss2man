import axios from 'axios';

const getHeader = tokenData => ({
  'access-token': tokenData.accessToken,
  'token-type': tokenData.tokenType,
  uid: tokenData.uid,
  client: tokenData.client,
});

const FavoriteService = {};

FavoriteService.handleFavorite = (roomId, token, action) => {
  const url = action === 'create' ? '/favorites/create' : '/favorites/delete';
  return axios.post(url, {room_id: roomId}, {headers: getHeader(token)});
};

FavoriteService.getFavoriteList = token => {
  return axios.get('/favorites', {headers: getHeader(token)});
};


export default FavoriteService;
