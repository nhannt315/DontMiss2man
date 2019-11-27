import axios from 'axios';

const BuildingService = {};

BuildingService.getBuildingList = (page, perPage) => {
  return axios.get(`/buildings?page=${page}&per_page=${perPage}`).then(res => res.data.data);
};

export default BuildingService;
