import React, {useRef} from 'react';
import {Layout, Button, Dropdown, Menu, Avatar} from 'antd';
import {connect} from 'react-redux';
import {Link} from 'react-router-dom';
import PropTypes from 'prop-types';
import scrollToComponent from 'react-scroll-to-component';
import i18n from '../../config/i18n';
import Logo from '../../assets/images/logo.png';
import './Layout.scss';
import {logout as logoutRedux} from '../../store/actions';
import ToTopButton from '../ToTopButton';

const {Header, Content} = Layout;

class MainLayout extends React.PureComponent {
  constructor(props) {
    super(props);
    this.firstElement = React.createRef();
    this.state = {
      showHeader: false,
      scrollPos: 0,
    };
  }

  componentDidMount() {
    window.addEventListener('scroll', this.handleScroll);
  }

  componentWillUnmount() {
    window.removeEventListener('scroll', this.handleScroll);
  }

  handleScroll = () => {
    const { scrollPos } = this.state;
    this.setState({
      scrollPos: document.body.getBoundingClientRect().top,
      showHeader: document.body.getBoundingClientRect().top > scrollPos
    });
  };

  toLoginPage = () => {
    const {history} = this.props;
    history.push('/login');
  };

  toTop = () => {
    scrollToComponent(this.firstElement.current);
  };

  render() {
    const {children, userData, logout, isAuthenticated} = this.props;
    const {showHeader} = this.state;
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
          <Header className={`header ${showHeader ? 'header--active' : 'header--hidden'}`}>
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
