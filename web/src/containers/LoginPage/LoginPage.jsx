import React, {PureComponent} from 'react';
import PropTypes from 'prop-types';
import {connect} from 'react-redux';
import LoginForm from '../../components/LoginForm';
import AuthBackground from '../../assets/images/bg.jpg';
import {login as loginRedux} from '../../store/actions';

class LoginPage extends PureComponent {

  constructor(props) {
    super(props);
    this.state = {
      leave: false,
    };
  }

  componentDidMount() {
    this.checkIsAuthenticated();
  }

  componentDidUpdate(prevProps) {
    this.checkIsAuthenticated();
  }

  checkIsAuthenticated = () => {
    const {history, isAuthenticated} = this.props;
    if (isAuthenticated) {
      this.setState({leave: true});
      setTimeout(() => history.push('/'), 700);
    }
  };

  handleSubmit = values => {
    const {login} = this.props;
    login(values.email, values.password, values.remember);
  };

  toRegisterPage = () => {
    const {history} = this.props;
    this.setState({leave: true});
    setTimeout(() => history.push('/registration'), 700);
  };

  render() {
    const {leave} = this.state;
    const {loading} = this.props;
    const formClassName = `animated ${leave ? 'bounceOutLeft' : 'bounceInRight'}`;
    return (
      <div
        className="auth_page"
        style={{backgroundImage: `linear-gradient(to right bottom, rgb(99, 125, 143), rgba(52, 52, 52, 0.6)), url(${AuthBackground})`}}>
        <div className="auth_box">
          <div className="auth_box-wrapper">
            <LoginForm
              className={formClassName} submit={this.handleSubmit}
              toRegisterPage={this.toRegisterPage} loading={loading}
            />
          </div>
        </div>
      </div>
    );
  }
}

LoginPage.propTypes = {
  history: PropTypes.object,
  isAuthenticated: PropTypes.bool,
  login: PropTypes.func,
  loading: PropTypes.bool,
};

LoginPage.defaultProps = {
  history: {},
  loading: false,
  isAuthenticated: false,
  login: () => {
  },
};

const mapStateToProps = state => ({
  isAuthenticated: state.auth.isAuthenticated,
  loading: state.auth.isProcessing,
});

const mapDispatchToProps = dispatch => ({
  login: (email, password, remember) => dispatch(loginRedux(email, password, remember)),
});

export default connect(mapStateToProps, mapDispatchToProps)(LoginPage);
