import React, {useState} from 'react';
import { Layout, Menu, Icon, Button } from 'antd';
import PropTypes from 'prop-types';

import i18n from '../../config/i18n';
import './Layout.scss';

const { Header, Sider, Content } = Layout;

const MainLayout = ({children}) => {

  const [collapsed, setCollapsed] = useState(false);

  const toggle = () => setCollapsed(!collapsed);

  return (
    <Layout id="app-bar">
      <Sider trigger={null} collapsible collapsed={collapsed}>
        <div className="logo" />
        <Menu theme="dark" mode="inline" defaultSelectedKeys={['1']}>
          <Menu.Item key="1">
            <Icon type="user" />
            <span>nav 1</span>
          </Menu.Item>
          <Menu.Item key="2">
            <Icon type="video-camera" />
            <span>nav 2</span>
          </Menu.Item>
          <Menu.Item key="3">
            <Icon type="upload" />
            <span>nav 3</span>
          </Menu.Item>
        </Menu>
      </Sider>
      <Layout>
        <Header className="header" style={{display: 'flex', flexDirection: 'row'}}>
          <div className="trigger">
            <Icon
              type={collapsed ? 'menu-unfold' : 'menu-fold'}
              onClick={toggle}
            />
          </div>
          <div className="vertical-align">
            <Button>{i18n.t('common.login')}</Button>
          </div>
        </Header>
        <Content
          className="main-content"
        >
          {children}
        </Content>
      </Layout>
    </Layout>
  );
};

MainLayout.propTypes = {
  children: PropTypes.node
};

MainLayout.defaultProps = {
  children: <React.Fragment> </React.Fragment>
};

export default MainLayout;
