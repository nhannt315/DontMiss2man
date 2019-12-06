import React, {PureComponent} from 'react';
import PropTypes from 'prop-types';
import './LoginPage.scss';
import LoginForm from '../../components/LoginForm';
import AuthBackground from '../../assets/images/bg.jpg';

class LoginPage extends PureComponent {

  constructor(props) {
    super(props);
    this.state = {
      leave: false,
    };
  }

  handleSubmit = values => {
    const {history} = this.props;
    this.setState({leave: true});
    setTimeout(() => history.push('/'), 700);
  };

  render() {
    const {leave} = this.state;
    const formClassName = `animated ${leave ? 'bounceOutLeft' : 'bounceInRight'}`;
    return (
      <div
        className="auth_page"
        style={{backgroundImage: `linear-gradient(to right bottom, rgb(99, 125, 143), rgba(52, 52, 52, 0.6)), url(${AuthBackground})`}}>
        <div className="auth_box">
          <div className="auth_box-wrapper">
            <LoginForm className={formClassName} submit={this.handleSubmit} />
          </div>
        </div>
      </div>
    );
  }
}

LoginPage.propTypes = {
  history: PropTypes.object,
};

LoginPage.defaultProps = {
  history: {},
};

export default LoginPage;
