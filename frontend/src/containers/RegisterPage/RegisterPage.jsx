import React, {PureComponent} from 'react';
import PropTypes from 'prop-types';
import {message} from 'antd';
import {connect} from 'react-redux';
import AuthBackground from '../../assets/images/bg.jpg';
import RegisterForm from '../../components/RegisterForm';
import {register as registerRedux} from '../../store/actions';
import i18n from '../../config/i18n';

class RegisterPage extends PureComponent {
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
    const {register} = this.props;
    register(values.email, values.password, values.password_confirm, this.registerSuccessCallback);
  };

  registerSuccessCallback = () => {
    const {history} = this.props;
    message.success(i18n.t('auth.pls_check_mail'), 6);
    this.setState({leave: true});
    setTimeout(() => history.push('/'), 700);
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
            <RegisterForm loading={loading} className={formClassName} submit={this.handleSubmit} />
          </div>
        </div>
      </div>
    );
  }
}

RegisterPage.propTypes = {
  history: PropTypes.object,
  isAuthenticated: PropTypes.bool,
  register: PropTypes.func,
  loading: PropTypes.bool,
};

RegisterPage.defaultProps = {
  history: {},
  isAuthenticated: false,
  loading: false,
  register: () => {
  },
};

const mapStateToProps = state => ({
  isAuthenticated: state.auth.isAuthenticated,
  loading: state.auth.isProcessing,
});

const mapDispatchToProps = dispatch => ({
  register: (email, password, passwordConfirm, callback) => dispatch(registerRedux(email, password, passwordConfirm, callback)),
});

export default connect(mapStateToProps, mapDispatchToProps)(RegisterPage);
