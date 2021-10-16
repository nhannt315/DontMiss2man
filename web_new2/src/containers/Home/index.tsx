import React, { useEffect, useState } from 'react';
import BuildingList from 'src/components/home/BuildingList';
import SearchDetail from 'src/components/home/SearchDetail';
import BuildingService from 'src/services/api/building';
import { IBuilding } from 'src/types/building';

const Home: React.FC = () => {
  const [buildingList, setBuildingList] = useState<IBuilding[]>([]);

  useEffect(() => {
    getBuildingList();
  }, []);

  const getBuildingList = async () => {
    try {
      const res = await BuildingService.getBuildingList(1, 10);
      setBuildingList(res.data.data.list);
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <div className="flex flex-row container mx-auto pt-4 bg-white mt-2 h-full px-6">
      <BuildingList className="w-8/12" list={buildingList} />
      <SearchDetail className="w-4/12" />
    </div>
  );
};

export default Home;
