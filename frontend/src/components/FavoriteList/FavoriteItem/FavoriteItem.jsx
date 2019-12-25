import React from 'react';
import PropTypes from 'prop-types';
import {Link} from 'react-router-dom';
import i18n from '../../../config/i18n';
import './FavoriteItem.scss';
import Image from '../../Image';
import CommonHelper from '../../../helpers/common';
import OverlayConfirmation from '../../OverlayConfirmation';

class FavoriteItem extends React.PureComponent {

  constructor(props) {
    super(props);
    this.state = {
      showDeleteConfirm: false,
      deleteLoading: false,
    };
  }

  handleDelete = () => {
    const {deleteFavorite, room} = this.props;
    this.setState({deleteLoading: true});
    deleteFavorite(room.id)
      .catch(() => this.setState({deleteLoading: false}));
  };

  render() {
    const {room} = this.props;
    const {showDeleteConfirm, deleteLoading} = this.state;
    return (
      <div className="favorite_item">
        <OverlayConfirmation
          show={showDeleteConfirm}
          title={i18n.t('common.favorite_delete_confirm')}
          okText={i18n.t('common.to_delete')}
          cancelText={i18n.t('common.cancel')}
          onOk={this.handleDelete}
          onCancel={() => this.setState({showDeleteConfirm: false})}
          loading={deleteLoading}
        />
        <div className="favorite_item__image">
          <Image src={room.building.photo_url} />
          <div className="favorite_item__delete_btn" onClick={() => this.setState({showDeleteConfirm: true})}>
            {i18n.t('common.delete_from_favorites')}
          </div>
        </div>
        <div className="favorite_item__info">
          <div className="favorite_item__name">
            <Link to={`/chintai/${room.id}`} target="_blank">
              {`${room.building.name} ${i18n.t('common.floor_number', {postProcess: 'interval', floor: room.floor})}`}
            </Link>
          </div>
          <div className="favorite_item__main_fee">
            {CommonHelper.convertYen(room.rent_fee)}
          </div>
          <div className="favorite_item__sub_fee">
          <span className="favorite_item__sub_fee__management">
            {CommonHelper.convertYen(room.management_cost)}
          </span>
            <span className="favorite_item__sub_fee__reikin">
            {CommonHelper.convertYen(room.reikin)}
          </span>
            <span className="favorite_item__sub_fee__shikikin">
            {CommonHelper.convertYen(room.shikikin)}
          </span>
          </div>
          <div className="favorite_item__extra">
            <span>{room.layout}</span>
            <span>{room.size}m<sup>2</sup></span>
            <span>{CommonHelper.getYearBuiltInJap(room.building.years_built)}</span>
            <span>{room.direction}</span>
          </div>
          <div className="favorite_item__access">
            <span>{room.building.access[0]}</span>
          </div>
        </div>
      </div>
    );
  }
}

FavoriteItem.propTypes = {
  room: PropTypes.object,
  deleteFavorite: PropTypes.func,
};

FavoriteItem.defaultProps = {
  room: {},
  deleteFavorite: () => {
  },
};

export default FavoriteItem;
