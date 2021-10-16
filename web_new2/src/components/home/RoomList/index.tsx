import React from 'react';
import { IRoom } from 'src/types/room';
import Link from 'next/link';
import useTranslation from 'next-translate/useTranslation';
import { convertYen } from 'src/utils/currency';
import styles from './style.module.scss';

interface Props {
  list: IRoom[];
}

const RoomList: React.FC<Props> = ({ list }) => {
  const { t, lang } = useTranslation('common');

  return (
    <div>
      <table className="w-full">
        <thead>
          <tr>
            <th className="text-gray-400 py-2 px-1 text-sm ">&nbsp;</th>
            <th className="text-gray-400 py-2 px-1 text-sm text-center">
              {t('floor')}
            </th>
            <th className="text-gray-400 py-2 px-1 text-sm text-center">
              {t('rent_fee')}
            </th>
            <th className="text-gray-400 py-2 px-1 text-sm text-center">
              {t('reikin')}/{t('shikikin')}
              {}
            </th>
            <th className="text-gray-400 py-2 px-1 text-sm text-center">
              {t('layout')}/{t('size')}
            </th>
            <th className="text-gray-400 py-2 px-1 text-sm text-center">
              {t('like')}
            </th>
            <th className="text-gray-400 py-2 px-1 text-sm ">&nbsp;</th>
          </tr>
        </thead>
        <tbody>
          {list.map((item) => (
            <tr
              className="w-8 h-32 py-2 border-b border-gray-300"
              key={item.id}
            >
              <td className="w-20 h-28">
                <img src={item.layout_image_url} alt={item.layout_image_url} />
              </td>
              <td className="text-center text-gray-500">
                {t('floor_number', {
                  postProcess: 'interval',
                  floor: item.floor,
                })}
              </td>
              <td className="text-center text-gray-500">
                <div>
                  <div className="room-rent-fee text-red-500 font-bold text-xl">
                    {convertYen(item.rent_fee, lang, t)}
                  </div>
                  <div>{convertYen(item.management_cost, lang, t)}</div>
                </div>
              </td>
              <td className="text-center text-gray-500">
                <div>
                  <div className={styles.reikin}>
                    {convertYen(item.reikin, lang, t)}
                  </div>
                  <div className={styles.shikikin}>
                    {convertYen(item.shikikin, lang, t)}
                  </div>
                </div>
              </td>
              <td className="text-center text-gray-500">
                <div>
                  <div>{item.layout}</div>
                  <div>
                    {item.size}m<sup>2</sup>
                  </div>
                </div>
              </td>
              <td>
                <div className="flex items-center justify-center">
                  <button className="border border-gray-400 text-gray-500 rounded text-xs py-2 font-bold px-4">
                    {t('add')}
                  </button>
                </div>
              </td>
              <td>
                <div className="text-center text-sm underline text-blue-600 font-bold">
                  <Link href="/">{t('see_detail')}</Link>
                </div>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default RoomList;
