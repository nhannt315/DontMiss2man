import React from 'react';
import {Layout, Button, Dropdown, Menu, Avatar} from 'antd';
import {connect} from 'react-redux';
import {Link} from 'react-router-dom';
import PropTypes from 'prop-types';
import i18n from '../../config/i18n';
import Logo from '../../assets/images/logo.png';
import './Layout.scss';
import {logout as logoutRedux} from '../../store/actions';

const {Header, Content} = Layout;

const MainLayout = ({children, history, userData, logout, isAuthenticated}) => {
  const toLoginPage = () => history.push('/login');
  const userMenu = (
    <Menu>
      <Menu.Item>
        <div>{i18n.t('common.my_page')}</div>
      </Menu.Item>
      <Menu.Item>
        <div onClick={logout}>{i18n.t('common.logout')}</div>
      </Menu.Item>
    </Menu>
  );
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
              <div className="logo-slogan">{i18n.t('common.slogan')}</div>
            </div>
            <div className="vertical-align">
              {isAuthenticated ? (
                <Dropdown overlay={userMenu}>
                  <div className="user_area">
                    <Avatar icon="user" />
                    <span>{userData.email}</span>
                  </div>
                </Dropdown>
              ) : (
                <Button onClick={toLoginPage}>{i18n.t('common.login')}</Button>
              )}
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
  userData: PropTypes.object,
  logout: PropTypes.func,
  isAuthenticated: PropTypes.bool,
};

MainLayout.defaultProps = {
  children: <React.Fragment> </React.Fragment>,
  history: {},
  userData: {},
  logout: () => {
  },
  isAuthenticated: false,
};

const mapStateToProps = state => ({
  isAuthenticated: state.auth.isAuthenticated,
  userData: state.auth.userData,
});

const mapDispatchToProps = dispatch => ({
  logout: () => dispatch(logoutRedux()),
});

export default connect(mapStateToProps, mapDispatchToProps)(MainLayout);
