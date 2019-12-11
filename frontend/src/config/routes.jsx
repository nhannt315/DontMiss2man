import React, {useEffect} from 'react';
import {Switch, Router, Route} from 'react-router';
import PropTypes from 'prop-types';
import {connect} from 'react-redux';
import Footer from '../components/Footer';
import {authCheckState as authCheckStateRedux} from '../store/actions';

import * as Containers from '../containers';

const buildRoutes = (history) => {
  return (
    <Router history={history}>
      <div className="outer-layout">
        <Switch>
          <Route exact path="/" component={Containers.HomePage} />
          <Route exact path="/chintai/:id" component={Containers.RoomDetailPage} />
          <Route exact path="/login" component={Containers.LoginPage} />
          <Route exact path="/registration" component={Containers.RegisterPage} />
          <Route exact path="/confirmation" component={Containers.AccountConfirmationPage} />
        </Switch>
        <Footer history={history} />
      </div>
    </Router>
  );
};

const RouterWrapper = ({history, authCheckState}) => {
  useEffect(() => {
    authCheckState();
  }, [authCheckState]);
  return <div className="router-wrapper">{buildRoutes(history)}</div>;
};

RouterWrapper.propTypes = {
  history: PropTypes.object,
  authCheckState: PropTypes.func,
};

RouterWrapper.defaultProps = {
  history: {},
  authCheckState: () => {
  },
};

const mapDispatchToProps = dispatch => ({
  authCheckState: () => dispatch(authCheckStateRedux()),
});


export default connect(null, mapDispatchToProps)(RouterWrapper);
