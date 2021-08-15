import React from 'react';

import './OverlayLoader.scss';

const OverlayLoader = () => {
  return (
    <div className="overlay-loader">
      <div className="overlay__inner">
        <div className="overlay__content"><span className="spinner" /></div>
      </div>
    </div>
  );
};

OverlayLoader.propTypes = {};

export default OverlayLoader;
