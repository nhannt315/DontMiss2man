import React from 'react';
import PropTypes from 'prop-types';
import './Marker.scss';

const Marker = ({type}) => {
  return (
    <div className={type === 'filled' ? 'pin2' : 'pin1'} />
  );
};

Marker.propTypes = {
  type: PropTypes.oneOf(['filled', 'outline']),
};

Marker.defaultProps = {
  type: 'filled',
};


export default Marker;
