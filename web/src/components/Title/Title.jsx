import React from 'react';
import PropTypes from 'prop-types';

import './Title.scss';

const Title = ({content, className}) => {
  return (
    <h2 className={[className, 'title'].join(' ')}>
      <span>{content}</span>
    </h2>
  );
};

Title.propTypes = {
  content: PropTypes.string,
  className: PropTypes.string,
};

Title.defaultProps = {
  content: '',
  className: '',
};

export default Title;
