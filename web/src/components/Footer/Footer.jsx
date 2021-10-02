import React from 'react';
import PropTypes from 'prop-types';
import moment from 'moment';
import './Footer.scss';

const Footer = ({history}) => {
  return (
    <div className="footer_layout">
      <div className="container">
        <div className="footer-content-center" onClick={() => history.push('/')}>
          <span>©{moment().year()} Dm2m.online</span>
        </div>
      </div>
    </div>
  );
};

Footer.propTypes = {
  history: PropTypes.object,
};

Footer.defaultProps = {
  history: {},
};

export default Footer;
