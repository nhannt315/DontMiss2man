import React, {useState} from 'react';
import PropTypes from 'prop-types';
import NotFoundImg from '../../assets/images/not_found.png';

const Image = ({className, styles, src, alt, fallbackUrl}) => {
  const [imgSrc, setSrc] = useState(src || NotFoundImg);
  return (
    <img className={className} style={styles} src={imgSrc} alt={alt} onError={() => setSrc(fallbackUrl)} />
  );
};

Image.propTypes = {
  src: PropTypes.string,
  alt: PropTypes.string.isRequired,
  fallbackUrl: PropTypes.string,
  className: PropTypes.string,
  styles: PropTypes.object,
};

Image.defaultProps = {
  src: null,
  fallbackUrl: NotFoundImg,
  className: '',
  styles: {},
};

export default Image;
