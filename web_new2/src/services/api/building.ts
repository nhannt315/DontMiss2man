import axios from 'axios';
import qs from 'query-string';
import { IBuilding } from 'src/types/building';
import { IResponse } from '../response';

export interface ISearchCondition {
  upper_fee: number;
  lower_fee: number;
  no_management_fee: boolean;
  no_reikin: boolean;
  no_shikikin: boolean;
  layout_types: string;
  building_type: string;
  upper_size: number;
  lower_size: number;
  years_built: number;
  with_furniture: boolean;
}

export interface IBuildingListResponse {
  list: IBuilding[];
  total_pages: number;
  total: number;
  page: number;
}

class BuildingService {
  getBuildingList(
    page: number,
    perPage: number,
    sort?: string,
    condition?: ISearchCondition
  ) {
    const url = `/buildings?page=${page}&per_page=${perPage}&sort=${sort}`;
    if (!condition) {
      return axios.get(url);
    }
    return axios.get<IResponse<IBuildingListResponse>>(
      `${url}&${qs.stringify(condition, { arrayFormat: 'bracket' })}`
    );
  }
}

export default new BuildingService();
