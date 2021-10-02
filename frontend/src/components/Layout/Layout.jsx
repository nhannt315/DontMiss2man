import React from 'react';
import {Layout, Button, Dropdown, Menu, Avatar, Modal} from 'antd';
import {connect} from 'react-redux';
import {Link} from 'react-router-dom';
import PropTypes from 'prop-types';
import throttleLodash from 'lodash/throttle';
import scrollToComponent from 'react-scroll-to-component';
import i18n from '../../config/i18n';
import Logo from '../../assets/images/logo.png';
import './Layout.scss';
import {
  logout as logoutRedux,
  removeUserFavorite as removeUserFavoriteRedux,
} from '../../store/actions';
import ToTopButton from '../ToTopButton';
import FavoriteList from '../FavoriteList';
import LanguageChanger from '../LanguageChanger';

const {Header, Content} = Layout;

class MainLayout extends React.PureComponent {
  constructor(props) {
    super(props);
    this.firstElement = React.createRef();
    this.state = {
      showHeader: true,
      scrollPos: 0,
      showFavoriteModal: false,
    };
  }

  componentDidMount() {
    window.addEventListener('scroll', throttleLodash(this.handleScroll, 100));
  }

  componentWillUnmount() {
    window.removeEventListener('scroll', throttleLodash(this.handleScroll, 100));
  }

  handleScroll = () => {
    const {scrollPos} = this.state;
    const topPos = document.body.getBoundingClientRect().top;
    if (topPos < -64 || topPos > scrollPos) {
      this.setState({
        scrollPos: topPos,
        showHeader: topPos > scrollPos,
      });
    }
  };

  toLoginPage = () => {
    const {history} = this.props;
    history.push('/login');
  };

  toTop = () => {
    scrollToComponent(this.firstElement.current);
  };

  render() {
    const {children, userData, logout, isAuthenticated, tokenData, removeUserFavorite, language} = this.props;
    const {showHeader, showFavoriteModal} = this.state;
    const userMenu = (
      <Menu>
        <Menu.Item onClick={() => this.setState({showFavoriteModal: true})}>
          <div>{i18n.t('common.like')}</div>
        </Menu.Item>
        <Menu.Item>
          <div onClick={logout}>{i18n.t('common.logout')}</div>
        </Menu.Item>
      </Menu>
    );

    return (
      <Layout id="app-bar">
        <Modal
          title={i18n.t('common.favorite_list')}
          visible={showFavoriteModal}
          onCancel={() => this.setState({showFavoriteModal: false})}
          footer={null}
          width="35vw"
        >
          <FavoriteList
            tokenData={tokenData} userFavoriteIds={userData.favorites}
            removeUserFavorite={removeUserFavorite} show={showFavoriteModal}
          />
        </Modal>
        <Layout>
          <Header className={`header ${showHeader ? 'header--active' : 'header--hidden'}`}>
            <div className="container" style={{display: 'flex', flexDirection: 'row'}}>
              <div className="logo">
                <Link to="/">
                  <img src={Logo} alt="Logo" />
                  <span>DM2M</span>
                </Link>
                <div className="logo-slogan">{i18n.t('common.slogan')}</div>
              </div>
              <LanguageChanger currentLng={language} />
              <div className="vertical-align">
                {isAuthenticated ? (
                  <Dropdown overlay={userMenu}>
                    <div className="user_area">
                      <Avatar icon="user" />
                      <span>{userData.email}</span>
                    </div>
                  </Dropdown>
                ) : (
                  <Button onClick={this.toLoginPage}>{i18n.t('common.login')}</Button>
                )}
              </div>
            </div>
          </Header>
          <Content className="container main-content" style={{marginTop: '1rem'}}>
            <div ref={this.firstElement} />
            {children}
          </Content>
          <div className="to_top_area">
            <ToTopButton onClick={this.toTop} />
          </div>
        </Layout>
      </Layout>
    );
  }
}


MainLayout.propTypes = {
  children: PropTypes.node,
  history: PropTypes.object,
  userData: PropTypes.object,
  logout: PropTypes.func,
  isAuthenticated: PropTypes.bool,
  tokenData: PropTypes.object,
  removeUserFavorite: PropTypes.func,
  language: PropTypes.string,
};

MainLayout.defaultProps = {
  children: <React.Fragment> </React.Fragment>,
  history: {},
  userData: {},
  logout: () => {
  },
  isAuthenticated: false,
  tokenData: {},
  removeUserFavorite: () => {
  },
  language: '',
};

const mapStateToProps = state => ({
  isAuthenticated: state.auth.isAuthenticated,
  userData: state.auth.userData,
  tokenData: state.auth.tokenData,
  language: state.ui.language,
});

const mapDispatchToProps = dispatch => ({
  logout: () => dispatch(logoutRedux()),
  removeUserFavorite: roomId => dispatch(removeUserFavoriteRedux(roomId)),
});

export default connect(mapStateToProps, mapDispatchToProps)(MainLayout);
