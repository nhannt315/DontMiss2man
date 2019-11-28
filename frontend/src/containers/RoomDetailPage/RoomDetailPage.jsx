import React, {useState, useEffect} from 'react';
import PropTypes from 'prop-types';
import {Carousel} from 'antd';
import {connect} from 'react-redux';
import i18n from '../../config/i18n';
import './RoomDetailPage.scss';
import {fetchRoom} from '../../store/actions';
import ListPlaceholder from '../../components/ListPlaceholder';

const RoomDetailPage = ({room, loading, error, fetchRoomDetail, match}) => {
  useEffect(() => {
    fetchRoomDetail(match.params.id);
  }, []);
  if (loading)
    return <ListPlaceholder itemCount={5} />;
  return (
    <div className="roomdetail">
      <div className="roomdetail-title">
        <h1>{room.building.name}</h1>
      </div>
      <div className="navigation-post">
        <ul>
          <li>{i18n.t('roomDetail.general_info')}</li>
          <li>{i18n.t('roomDetail.images')}</li>
          <li>{i18n.t('roomDetail.detail_info')}</li>
          <li>{i18n.t('roomDetail.map')}</li>
        </ul>
      </div>
      <div className="roomdetail-image-list">
        <Carousel>
          {room.images.map(image => (
            <div key={image.id} className="carousel-image-wrapper">
              <img src={image.url} alt={image.description} />
            </div>
          ))}
        </Carousel>
      </div>
      <div className="roomdetail-general-info">
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
              <td>{room.size}</td>
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
    </div>
  );
};

RoomDetailPage.propTypes = {
  room: PropTypes.object,
  loading: PropTypes.bool,
  error: PropTypes.object,
  fetchRoomDetail: PropTypes.func,
  match: PropTypes.object,
};

RoomDetailPage.defaultProps = {
  room: {},
  loading: false,
  error: null,
  fetchRoomDetail: () => {
  },
  match: {},
};

const mapStateToProps = state => ({
  room: state.room.room,
  loading: state.room.loading,
  error: state.room.error,
});

const mapDispatchToProps = dispatch => ({
  fetchRoomDetail: roomId => dispatch(fetchRoom(roomId)),
});

export default connect(mapStateToProps, mapDispatchToProps)(RoomDetailPage);
