import { IRoom } from './room';

export interface IBuilding {
  id: number;
  name: string;
  address: string;
  access: string[];
  year_built: Date;
  building_type: string;
  structure: string;
  storeys: number;
  underground_storeys: number;
  photo_url: string;
  longitude: number;
  latitude: number;
  average_size: number;
  average_fee: number;
  distance: number;
  condition_type: number;
  office_id: number;
  rooms: IRoom[];
  created_at: Date;
  updated_at: Date;
}
