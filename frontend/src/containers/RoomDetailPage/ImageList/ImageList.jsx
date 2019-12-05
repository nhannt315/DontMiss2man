import React from 'react';
import PropTypes from 'prop-types';
import ImageSlider from '../../../components/ImageSlider';
import './ImageList.scss';
import i18n from '../../../config/i18n';
import Title from '../../../components/Title';

const ImageList = ({room}) => {
  return (
    <div>
      <ImageSlider images={room.images} />
    </div>
  )
};

ImageList.propTypes = {
  room: PropTypes.object,
};

ImageList.defaultProps = {
  room: {},
};

export default ImageList;
