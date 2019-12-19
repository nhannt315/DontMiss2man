import React, {useState} from 'react';
import PropTypes from 'prop-types';
import {Button, Icon, message} from 'antd';
import i18n from '../../../config/i18n';
import './RoomItem.scss';
import CommonHelper from '../../../helpers/common';
import Image from '../../Image';

const RoomItem = ({room, history, userData, handleFavoriteAction, isAuthenticated}) => {
  const [loading, setLoading] = useState(false);
  const handleDetailClicked = () => {
    history.push(`/chintai/${room.id}`);
  };
  const isFavorited = isAuthenticated && (userData.favorites.includes(room.id));
  const handleFavoriteButtonClicked = () => {
    if (!isAuthenticated){
      message.warning(i18n.t('common.need_login'));
      return;
    }
    const action = isFavorited ? 'delete' : 'create';
    setLoading(true);
    handleFavoriteAction(room.id, action)
      .then(() => setLoading(false));
  };
  return (
    <tr className="roomitem">
      <td>
        <Image className="room-image" src={room.layout_image_url} alt={room.layout_image_url} />
      </td>
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
          <div>{room.size}m<sup>2</sup></div>
        </div>
      </td>
      <td>
        <div>
          <Button loading={loading} onClick={handleFavoriteButtonClicked}>
            {!loading && <Icon type="heart" theme={isFavorited ? 'filled' : 'outlined'} />}
            {i18n.t('common.add')}
          </Button>
        </div>
      </td>
      <td>
        <div className="roomitem-detail-link" onClick={handleDetailClicked}>
          {i18n.t('common.see_detail')}
        </div>
      </td>
    </tr>
  );
};

RoomItem.propTypes = {
  room: PropTypes.object,
  history: PropTypes.object.isRequired,
  userData: PropTypes.object,
  handleFavoriteAction: PropTypes.func.isRequired,
  isAuthenticated: PropTypes.bool,
};

RoomItem.defaultProps = {
  room: {},
  userData: {},
  isAuthenticated: false,
};

export default RoomItem;
