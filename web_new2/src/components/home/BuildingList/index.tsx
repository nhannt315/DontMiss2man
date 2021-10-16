import React from 'react';
import { IBuilding } from 'src/types/building';
import BuildingItem from '../BuildingItem';

interface IProps {
  className?: string;
  list: IBuilding[];
}

const BuildingList: React.FC<IProps> = ({ className, list }) => {
  return (
    <div className={className}>
      {list?.map((item) => (
        <BuildingItem key={item.id} item={item} />
      ))}
    </div>
  );
};

export default BuildingList;
