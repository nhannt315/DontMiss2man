import React, {useEffect} from 'react';
import {Button} from 'antd';
import PropTypes from 'prop-types';
import {connect} from 'react-redux';
import Layout from '../../components/Layout';
import i18n from '../../config/i18n';
import './AccountConfirmationPage.scss';

const AccountConfirmationPage = ({history, isAuthenticated}) => {
  useEffect(() => {
    if (isAuthenticated)
      history.push('/');
  }, [isAuthenticated, history]);
  return (
    <Layout history={history}>
      <div className="confirmation_wrapper">
        <div className="confirmation">
          <h3 className="confirmation-title">
            {i18n.t('confirmation.mail_confirm_complete')}
          </h3>
          <p className="confirmation-content">
            {i18n.t('confirmation.please_login')}
          </p>
          <div className="confirmation-button_wrapper">
            <Button
              type="primary" size="large"
              className="confirmation-button"
              onClick={() => history.push('/login')}
            >
              {i18n.t('confirmation.to_login_page')}
            </Button>
          </div>
        </div>
      </div>
    </Layout>
  );
};

AccountConfirmationPage.propTypes = {
  history: PropTypes.object,
  isAuthenticated: PropTypes.bool,
};

AccountConfirmationPage.defaultProps = {
  history: {},
  isAuthenticated: false,
};

const mapStateToProps = state => ({
  isAuthenticated: state.auth.isAuthenticated,
});

export default connect(mapStateToProps)(AccountConfirmationPage);
