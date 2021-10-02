import React from 'react';
import PropTypes from 'prop-types';

import './ImageSlide.scss';

const ImageSlide = ({url}) => {
  return (
    <div className="image-slide">
      <img src={url} alt={url} />
    </div>
  );
};

ImageSlide.propTypes = {
  url: PropTypes.string.isRequired,
};

export default ImageSlide;
