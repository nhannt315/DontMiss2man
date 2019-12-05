import React from 'react';
import PropTypes from 'prop-types';
import i18n from '../../../config/i18n';
import './PropertyInfo.scss';
import CommonHelper from '../../../helpers/common';

const PropertyInfo = ({room}) => {
  return (
    <div className="property_view_note">
      <div className="property_view_note-info">
        <div className="property_view_note-list">
          <span className="property_view_note-emphasis">{CommonHelper.convertYen(room.rent_fee)}</span>
          <span>{i18n.t('roomDetail.management_fee', {fee: CommonHelper.convertYen(room.management_cost)})}</span>
        </div>
        <div className="property_view_note-list">
          <span>{i18n.t('roomDetail.reikin', {fee: CommonHelper.convertYen(room.reikin)})}</span>
          <span>{i18n.t('roomDetail.shikikin', {fee: CommonHelper.convertYen(room.shikikin)})}</span>
        </div>
      </div>
    </div>
  );
};

PropertyInfo.propTypes = {
  room: PropTypes.object,
};

PropertyInfo.defaultProps = {
  room: {},
};

export default PropertyInfo;
