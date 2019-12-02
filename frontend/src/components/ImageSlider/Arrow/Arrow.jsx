import React from 'react';
import PropTypes from 'prop-types';

import './Arrow.scss';

const Arrow = ({direction, clickFunction, glyph}) => {
  return (
    <div
      className={`slide-arrow ${direction}`}
      onClick={clickFunction}
    >
      {glyph}
    </div>
  );
};

Arrow.propTypes = {
  direction: PropTypes.oneOf(['left', 'right']),
  clickFunction: PropTypes.func,
  glyph: PropTypes.string,
};

Arrow.defaultProps = {
  direction: 'left',
  clickFunction: () => {
  },
  glyph: '',
};

export default Arrow;
