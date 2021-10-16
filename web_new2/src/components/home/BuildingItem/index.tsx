import React from 'react';
import { IBuilding } from 'src/types/building';
import useTranslation from 'next-translate/useTranslation';
import dayjs from 'dayjs';
import RoomList from '../RoomList';

interface IProps {
  item: IBuilding;
}

const BuildingItem: React.FC<IProps> = ({ item }) => {
  const { t } = useTranslation('common');

  const getYearBuildInJap = (src: string) => {
    const years = dayjs().year() - dayjs(src).year();
    if (years > 2) {
      return t('number_of_year_built', { years: years });
    }
    return t('newly_built');
  };
  return (
    <div className="mb-4">
      <div className="h-1 bg-gray-400" />
      <div className="flex flex-row">
        <div className="pr-4 pt-2">
          <img alt={item.name} src={item.photo_url} width={120} height={240} />
        </div>
        <div className="flex-1 flex flex-col pt-1 ml-6">
          <div>
            <span className="border border-green-300 text-xs text-green-300 text-center p-1">
              賃貸マンション
            </span>
          </div>
          <div className="font-bold text-xl text-gray-500 pt-2">
            {item.name}
          </div>
          <div className="flex flex-row divide-dotted divide-x-2 divide-gray-300 border-dotted border-2 text-xs space-x-4 border-r-0 border-l-0">
            <div className="pl-2 pt-2 w-3/12">{item.address}</div>
            <div className="pl-2 py-2 w-7/12">
              {item.access.map((access) => {
                return <div key={access}>{access}</div>;
              })}
            </div>
            <div className="pl-2 pt-2 w-2/12">
              <div>{getYearBuildInJap(item.year_built.toString())}</div>
              <div>
                {t('number_of_storeys', {
                  storeys: item.storeys + item.underground_storeys,
                })}
              </div>
            </div>
          </div>
        </div>
      </div>
      <RoomList list={item.rooms} />
    </div>
  );
};

export default BuildingItem;
