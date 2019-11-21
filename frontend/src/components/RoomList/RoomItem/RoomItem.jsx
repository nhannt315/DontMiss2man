import React from 'react';
import PropTypes from 'prop-types';
import {Button} from 'antd';
import i18n from '../../../config/i18n';
import './RoomItem.scss';
import CommonHelper from '../../../helpers/common';

const RoomItem = ({room}) => {
  return (
    <tr className="roomitem">
      <td><img className="room-image" src={room.images[0].url} alt={room.images[0].url} /></td>
      <td>{i18n.t('common.floor_number', {floor: room.floor})}</td>
      <td>
        <div>
          <div className="room-rent-fee">{CommonHelper.convertYen(room.rent_fee)}</div>
          <div>{CommonHelper.convertYen(room.management_cost)}</div>
        </div>
      </td>
      <td>
        <div>
          <div className="roomitem-reikin">{CommonHelper.convertYen(room.reikin)}</div>
          <div className="roomitem-shikikin">{CommonHelper.convertYen(room.shikikin)}</div>
        </div>
      </td>
      <td>
        <div>
          <div>{room.layout}</div>
          <div>{room.size}m2</div>
        </div>
      </td>
      <td>
        <div>
          <Button icon="heart">{i18n.t('common.add')}</Button>
        </div>
      </td>
    </tr>
  );
};

RoomItem.propTypes = {
  room: PropTypes.object,
};

RoomItem.defaultProps = {
  room: {},
};

export default RoomItem;
