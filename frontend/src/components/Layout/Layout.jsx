import React from 'react';
import {Layout, Button} from 'antd';
import {Link} from 'react-router-dom';
import PropTypes from 'prop-types';

import i18n from '../../config/i18n';
import Logo from '../../assets/images/logo.png';
import './Layout.scss';

const {Header, Content} = Layout;

const MainLayout = ({children, history}) => {
  const toLoginPage = () => history.push('/login');
  return (
    <Layout id="app-bar">
      <Layout>
        <Header className="header">
          <div className="container" style={{display: 'flex', flexDirection: 'row'}}>
            <div className="logo">
              <Link to="/">
                <img src={Logo} alt="Logo" />
                <span>DM2M</span>
              </Link>
            </div>
            <div className="vertical-align">
              <Button onClick={toLoginPage}>{i18n.t('common.login')}</Button>
            </div>
          </div>
        </Header>
        <Content className="container main-content" style={{marginTop: '1rem'}}>
          {children}
        </Content>
      </Layout>
    </Layout>
  );
};

MainLayout.propTypes = {
  children: PropTypes.node,
  history: PropTypes.object,
};

MainLayout.defaultProps = {
  children: <React.Fragment> </React.Fragment>,
  history: {},
};

export default MainLayout;
