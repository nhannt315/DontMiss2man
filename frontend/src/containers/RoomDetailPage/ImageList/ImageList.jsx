import React from 'react';
import PropTypes from 'prop-types';

import './ImageList.scss';
import {Carousel} from 'antd';

const ImageList = ({room}) => {
  return (
    <div className="roomdetail-image-list">
      <Carousel>
        {room.images.map(image => (
          <div key={image.id} className="carousel-image-wrapper">
            <div style={{
              backgroundImage: `url(${image.url})`,
              backgroundSize: 'cover',
              backgroundPosition: 'center',
            }} />
          </div>
        ))}
      </Carousel>
    </div>
  );
};

ImageList.propTypes = {
  room: PropTypes.object,
};

ImageList.defaultProps = {
  room: {},
};

export default ImageList;
