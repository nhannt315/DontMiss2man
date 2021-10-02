import React, {useState} from 'react';
import PropTypes from 'prop-types';
import Arrow from './Arrow';
import ImageSlide from './ImageSlide';

import './ImageSlider.scss';

const ImageSlider = ({images, className, style}) => {
  const [currentImage, setCurrent] = useState(0);
  const previousSlide = () => {
    const lastIndex = images.length - 1;
    const shouldResetIndex = currentImage === 0;
    const index = shouldResetIndex ? lastIndex : currentImage - 1;
    setCurrent(index);
  };

  const nextSlide = () => {
    const lastIndex = images.length - 1;
    const shouldResetIndex = currentImage === lastIndex;
    const index = shouldResetIndex ? 0 : currentImage + 1;
    setCurrent(index);
  };
  return (
    <div className={['image-slider', className].join(' ')} style={style}>
      <div className="image-list">
        <Arrow direction="left" clickFunction={previousSlide} glyph="&#9664;" />
        <ImageSlide url={images[currentImage].url} />
        <Arrow direction="right" clickFunction={nextSlide} glyph="&#9654;" />
      </div>
      <div className="image_description">
        <div className="image_description-caption">{images[currentImage].description}</div>
        <div className="image_description-caption_number">{currentImage + 1}/{images.length}</div>
      </div>
    </div>
  );
};

ImageSlider.propTypes = {
  images: PropTypes.array,
  className: PropTypes.string,
  style: PropTypes.object,
};

ImageSlider.defaultProps = {
  images: [],
  className: '',
  style: {},
};

export default ImageSlider;
