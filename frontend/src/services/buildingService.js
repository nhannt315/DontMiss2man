import axios from 'axios';

const BuildingService = {};

BuildingService.getBuildingList = (page, perPage, sort, condition) => {
  return axios.get(`/buildings?page=${page}&per_page=${perPage}&sort=${sort}`).then(res => res.data.data);
};

export default BuildingService;
