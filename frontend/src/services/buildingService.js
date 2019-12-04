import axios from 'axios';
import qs from 'query-string';

const BuildingService = {};

BuildingService.getBuildingList = (page, perPage, sort, condition) => {
  const url = `/buildings?page=${page}&per_page=${perPage}&sort=${sort}`;
  if (!condition) {
    return axios.get(url).then(res => res.data.data);
  }
  const searchCondition = {
    upper_fee: condition.rentFee.upper,
    lower_fee: condition.rentFee.lower,
    no_management_fee: condition.rentFee.noManagementFee,
    no_reikin: condition.rentFee.noReikin,
    no_shikikin: condition.rentFee.noShikikin,
    layout_types: condition.layoutType,
    building_type: condition.buildingType,
    upper_size: condition.size.upper,
    lower_size: condition.size.lower,
    years_built: condition.years_built,
  };
  return axios.get(`${url}&${qs.stringify(searchCondition, {arrayFormat: 'bracket'})}`)
    .then(res => res.data.data);
};

export default BuildingService;
