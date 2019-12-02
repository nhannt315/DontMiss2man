import React from 'react';
import PropTypes from 'prop-types';
import Title from '../../../components/Title';
import i18n from '../../../config/i18n';

const GenerailInfo = ({room, ref}) => {
  return (
    <div className="roomdetail-general-info">
      <Title content={i18n.t('roomDetail.general_info')}/>
      <table className="table table-bordered">
        <tbody>
          <tr>
            <td className="title">{i18n.t('roomDetail.current_address')}</td>
            <td colSpan={3}>{room.building.address}</td>
          </tr>
          <tr>
            <td className="title">{i18n.t('roomDetail.walk_to_station')}</td>
            <td colSpan={3}>
              {room.building.access.map(ele => <div key={ele}>{ele}</div>)}
            </td>
          </tr>
          <tr>
            <td className="title">{i18n.t('roomDetail.layout')}</td>
            <td>{room.layout}</td>
            <td className="title">{i18n.t('roomDetail.size')}</td>
            <td>{room.size}m2</td>
          </tr>
          <tr>
            <td className="title">{i18n.t('roomDetail.year_built')}</td>
            <td>{room.building.year_built}</td>
            <td className="title">{i18n.t('roomDetail.floor')}</td>
            <td>{room.floor}{i18n.t('roomDetail.floor')}</td>
          </tr>
          <tr>
            <td className="title">{i18n.t('roomDetail.direction')}</td>
            <td>{room.direction}</td>
          </tr>
        </tbody>
      </table>
    </div>
  );
};

GenerailInfo.propTypes = {
  room: PropTypes.object,
};

GenerailInfo.defaultProps = {
  room: {},
};

export default GenerailInfo;
