import React from 'react';
import PropTypes from 'prop-types';

import './ImageSlide.scss';

const ImageSlide = ({url}) => {
  const styles = {
    backgroundImage: `url(${url})`,
    // backgroundSize: 'cover',
    backgroundPosition: 'center',
    backgroundRepeat: 'no-repeat',
  };

  return (
    <div className="image-slide" style={styles} />
  );
};

ImageSlide.propTypes = {
  url: PropTypes.string.isRequired,
};

export default ImageSlide;
