import React, {useRef, useEffect} from 'react';
import PropTypes from 'prop-types';
import {Button, Col, Row} from 'antd';
import {connect} from 'react-redux';
import i18n from '../../config/i18n';
import './RoomDetailPage.scss';
import {fetchRoom} from '../../store/actions';
import Layout from '../../components/Layout';
import ListPlaceholder from '../../components/ListPlaceholder';
import Title from '../../components/Title';
import GeneralInfo from './Generalnfo';
import DetailInfo from './DetailInfo';
import AgentInfo from './AgentInfo';
import MapInfo from './MapInfo';
import ImageList from './ImageList';
import PropertyInfo from './PropertyInfo';

const RoomDetailPage = ({room, loading, error, fetchRoomDetail, match, history}) => {
  const imageElement = useRef(null);
  const generalInfoElement = useRef(null);
  const facilityElement = useRef(null);
  const detailInfoElement = useRef(null);
  const mapElement = useRef(null);
  useEffect(() => {
    fetchRoomDetail(match.params.id);
  }, [fetchRoomDetail, match]);
  if (loading)
    return <ListPlaceholder itemCount={5} />;
  return (
    <Layout history={history}>
      <div className="roomdetail">
        <div className="roomdetail-title">
          <h1>{room.building.name}</h1>
          <div>
            <Button icon="heart">{i18n.t('common.add')}</Button>
          </div>
        </div>
        <Row>
          <Col span={24}>
            <PropertyInfo room={room} />
            <div ref={imageElement}>
              <ImageList room={room} />
            </div>
            <div ref={generalInfoElement}>
              <GeneralInfo room={room} />
            </div>
            <div ref={facilityElement}>
              <Title content={i18n.t('roomDetail.facilities')} />
              <div className="roomdetail-facilities">
                {room.facilities}
              </div>
            </div>
            <div ref={detailInfoElement}>
              <DetailInfo room={room} />
            </div>
            <div ref={mapElement}>
              <MapInfo latitude={room.building.latitude} longitude={room.building.longitude} language="ja" />
            </div>
            <div>
              <AgentInfo room={room} />
            </div>
          </Col>
        </Row>
      </div>
    </Layout>
  );
};

RoomDetailPage.propTypes = {
  room: PropTypes.object,
  loading: PropTypes.bool,
  error: PropTypes.object,
  fetchRoomDetail: PropTypes.func,
  match: PropTypes.object,
  history: PropTypes.object,
};

RoomDetailPage.defaultProps = {
  room: {},
  loading: false,
  error: null,
  fetchRoomDetail: () => {
  },
  match: {},
  history: {},
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
