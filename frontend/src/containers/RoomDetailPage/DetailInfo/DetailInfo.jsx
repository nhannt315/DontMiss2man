import React from 'react';
import PropTypes from 'prop-types';
import moment from 'moment';
import Title from '../../../components/Title';
import i18n from '../../../config/i18n';

const DetailInfo = ({room}) => {
  return (
    <div className="roomdetail-detail-info">
      <Title content={i18n.t('roomDetail.building_info')} />
      <table className="table table-bordered">
        <tbody>
          <tr>
            <td className="title">{i18n.t('roomDetail.detail_layout')}</td>
            <td>{room.layout_detail}</td>
            <td className="title">{i18n.t('roomDetail.structure')}</td>
            <td>{room.building.structure}</td>
          </tr>
          <tr>
            <td className="title">{i18n.t('roomDetail.number_of_floor')}</td>
            <td>{room.floor}/{room.building.storeys}</td>
            <td className="title">{i18n.t('roomDetail.time_built')}</td>
            <td>{moment(room.building.year_built).format('YYYY年MM月')}</td>
          </tr>
          <tr>
            <td className="title">{i18n.t('roomDetail.move_in')}</td>
            <td>{room.move_in}</td>
            <td className="title">{i18n.t('roomDetail.car_parking')}</td>
            <td>{room.car_park}</td>
          </tr>
          <tr>
            <td className="title">{i18n.t('roomDetail.condition')}</td>
            <td>{room.condition}</td>
            <td className="title">{i18n.t('roomDetail.deal_type')}</td>
            <td>{room.deal_type}</td>
          </tr>
          <tr>
            <td className="title">{i18n.t('roomDetail.other_fees')}</td>
            <td colSpan={3}>{room.other_fees}</td>
          </tr>
          <tr>
            <td className="title">{i18n.t('roomDetail.note')}</td>
            <td colSpan={3}>{room.note}</td>
          </tr>
        </tbody>
      </table>
    </div>
  );
};

DetailInfo.propTypes = {
  room: PropTypes.object,
};

DetailInfo.defaultProps = {
  room: {},
};

export default DetailInfo;
