import React from 'react';
import PropTypes from 'prop-types';
import {Icon} from 'antd';

import './ToTopButton.scss';

const ToTopButton = ({onClick, className, style}) => {
  return (
    <div onClick={onClick} className={['to_top_btn ripple', className].join(' ')} style={style}>
      <Icon type="up" />
    </div>
  );
};

ToTopButton.propTypes = {
  onClick: PropTypes.func,
  className: PropTypes.string,
  style: PropTypes.object,
};

ToTopButton.defaultProps = {
  onClick: () => {
  },
  className: '',
  style: {},
};

export default ToTopButton;
