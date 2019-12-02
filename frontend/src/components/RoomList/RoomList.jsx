import React from 'react';
import PropTypes from 'prop-types';

import './RoomList.scss';
import i18n from '../../config/i18n';
import RoomItem from './RoomItem';

const RoomList = ({list, history}) => {
  return (
    <table className="room-list">
      <thead>
        <tr>
          <th>&nbsp;</th>
          <th className="text-center">{i18n.t('common.floor')}</th>
          <th className="text-center">{i18n.t('common.rent_fee')}</th>
          <th className="text-center">{i18n.t('common.reikin')}/{i18n.t('common.shikikin')}{}</th>
          <th className="text-center">{i18n.t('common.layout')}/{i18n.t('common.size')}</th>
          <th className="text-center">{i18n.t('common.like')}</th>
          <th>&nbsp;</th>
        </tr>
      </thead>
      <tbody>
        {list.map(room => <RoomItem key={room.id} history={history} room={room} />)}
      </tbody>
    </table>
  );
};

RoomList.propTypes = {
  list: PropTypes.array,
  history: PropTypes.object.isRequired,
};

RoomList.defaultProps = {
  list: [],
};

export default RoomList;
