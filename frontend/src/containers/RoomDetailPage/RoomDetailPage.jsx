import React, {useRef, useEffect, useState} from 'react';
import scrollToComponent from 'react-scroll-to-component';
import PropTypes from 'prop-types';
import {Button, Col, Row, message, Icon} from 'antd';
import {connect} from 'react-redux';
import i18n from '../../config/i18n';
import './RoomDetailPage.scss';
import * as actions from '../../store/actions';
import FavoriteService from '../../services/favoriteService';
import Layout from '../../components/Layout';
import ListPlaceholder from '../../components/ListPlaceholder';
import Title from '../../components/Title';
import GeneralInfo from './Generalnfo';
import DetailInfo from './DetailInfo';
import AgentInfo from './AgentInfo';
import MapInfo from './MapInfo';
import ImageList from './ImageList';
import PropertyInfo from './PropertyInfo';

const RoomDetailPage = props => {
  const {
    room, loading, error, fetchRoomDetail, match, history,
    addUserFavorite, removeUserFavorite, userData, tokenData, isAuthenticated,
  } = props;
  const [favoriteLoading, setFavoriteLoading] = useState(false);
  const [isInitialized, setInitialize] = useState(false);
  const firstElement = useRef(null);
  useEffect(() => {
    fetchRoomDetail(match.params.id);
    if (!isInitialized) {
      scrollToComponent(firstElement.current);
      setInitialize(true);
    }
  }, [fetchRoomDetail, match, isInitialized]);
  const isFavorited = isAuthenticated && (userData.favorites.includes(room.id));
  const handleFavoriteButton = () => {
    if (!isAuthenticated) {
      message.warning(i18n.t('common.need_login'));
      return;
    }
    const action = isFavorited ? 'delete' : 'create';
    setFavoriteLoading(true);
    FavoriteService.handleFavorite(room.id, tokenData, action)
      .then(() => {
        setFavoriteLoading(false);
        if (action === 'create')
          addUserFavorite(room.id);
        else if (action === 'delete')
          removeUserFavorite(room.id);
      });
  };

  if (loading)
    return <ListPlaceholder itemCount={5} />;
  return (
    <Layout history={history}>
      <div className="roomdetail">
        <div ref={firstElement} />
        <div className="roomdetail-title">
          <h1>{room.building.name}</h1>
          <div>
            <Button loading={favoriteLoading} onClick={handleFavoriteButton}>
              {!favoriteLoading && <Icon type="heart" theme={isFavorited ? 'filled' : 'outlined'} />}
              {i18n.t('common.add')}
            </Button>
          </div>
        </div>
        <Row>
          <Col span={24}>
            <PropertyInfo room={room} />
            <div>
              <ImageList room={room} />
            </div>
            <div>
              <GeneralInfo room={room} />
            </div>
            <div>
              <Title content={i18n.t('roomDetail.facilities')} />
              <div className="roomdetail-facilities">
                {room.facilities}
              </div>
            </div>
            <div>
              <DetailInfo room={room} />
            </div>
            <div>
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
  addUserFavorite: PropTypes.func,
  removeUserFavorite: PropTypes.func,
  userData: PropTypes.object,
  tokenData: PropTypes.object,
  isAuthenticated: PropTypes.bool,
};

RoomDetailPage.defaultProps = {
  room: {},
  loading: false,
  error: null,
  fetchRoomDetail: () => {
  },
  match: {},
  history: {},
  addUserFavorite: () => {
  },
  removeUserFavorite: () => {
  },
  userData: {},
  tokenData: {},
  isAuthenticated: false,
};

const mapStateToProps = state => ({
  room: state.room.room,
  loading: state.room.loading,
  error: state.room.error,
  userData: state.auth.userData,
  tokenData: state.auth.tokenData,
  isAuthenticated: state.auth.isAuthenticated,
});

const mapDispatchToProps = dispatch => ({
  fetchRoomDetail: roomId => dispatch(actions.fetchRoom(roomId)),
  addUserFavorite: roomId => dispatch(actions.addUserFavorite(roomId)),
  removeUserFavorite: roomId => dispatch(actions.removeUserFavorite(roomId)),
});

export default connect(mapStateToProps, mapDispatchToProps)(RoomDetailPage);
