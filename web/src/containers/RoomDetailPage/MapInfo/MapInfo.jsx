import React from 'react';
import PropTypes from 'prop-types';
import {Map, Marker, GoogleApiWrapper} from 'google-maps-react';
import './MapInfo.scss';

const MapInfo = ({latitude, longitude, google, language}) => {
  return (
    <div className="map-wrapper">
      <Map google={google} style={{height: '30rem', width: '100%'}}
           initialCenter={{lat: latitude, lng: longitude}} alt={language}>
        <Marker />
      </Map>
    </div>
  );
};

MapInfo.propTypes = {
  latitude: PropTypes.number,
  longitude: PropTypes.number,
  google: PropTypes.object,
  language: PropTypes.string,
};

MapInfo.defaultProps = {
  latitude: 0,
  longitude: 0,
  google: {},
  language: 'ja',
};

export default GoogleApiWrapper(props => ({
  apiKey: (process.env.REACT_APP_GG_MAP_API_KEY),
  language: props.language,
}))(MapInfo);
