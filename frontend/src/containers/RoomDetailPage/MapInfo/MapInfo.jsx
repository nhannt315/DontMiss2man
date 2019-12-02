import React from 'react';
import PropTypes from 'prop-types';
import GoogleMapReact from 'google-map-react';
import Title from '../../../components/Title';
import i18n from '../../../config/i18n';
import Marker from '../../../components/Marker';

const MapInfo = ({room}) => {
  return (
    <React.Fragment>
      <Title content={i18n.t('roomDetail.map')} />
      <div style={{ height: '200px', width: '100%' }}>
        <GoogleMapReact
          bootstrapURLKeys={{ key: 'AIzaSyDMRiIuhVoM64aSk2gAyhyWhuOGNGBMHWU'}}
          defaultCenter={{lat: room.building.latitude, lng: room.building.longitude}}
          defaultZoom={14}
        >
          <Marker
            lat={room.building.latitude}
            lng={room.building.longitude}
            type="outline"
          />
        </GoogleMapReact>
      </div>
    </React.Fragment>
  );
};

MapInfo.propTypes = {
  room: PropTypes.object,
};

MapInfo.defaultProps = {
  room: {},
};

export default MapInfo;
