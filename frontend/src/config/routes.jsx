import React, {useEffect, useState} from 'react';
import {Switch, Router, Route} from 'react-router';
import PropTypes from 'prop-types';
import {connect} from 'react-redux';
import Footer from '../components/Footer';
import {
  authCheckState as authCheckStateRedux,
  changeLanguage as changeLanguageRedux,
} from '../store/actions';

import * as Containers from '../containers';
import {LOCALE_KEY} from '../constants/key';
import i18n from './i18n';
import moment from 'moment';

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

const RouterWrapper = ({history, authCheckState, changeLanguage, currentLanguage}) => {
  useEffect(() => {
    authCheckState();
    const language = localStorage.getItem(LOCALE_KEY) || 'ja';
    i18n.changeLanguage(language);
    changeLanguage(language);
  }, []);
  i18n.on('languageChanged', lng => {
    if (lng !== currentLanguage){
      moment.locale(lng);
      localStorage.setItem(LOCALE_KEY, lng);
      changeLanguage(lng);
    }
  });
  return <div className="router-wrapper">{buildRoutes(history)}</div>;
};

RouterWrapper.propTypes = {
  history: PropTypes.object,
  authCheckState: PropTypes.func,
  changeLanguage: PropTypes.func,
  currentLanguage: PropTypes.string,
};

RouterWrapper.defaultProps = {
  history: {},
  authCheckState: () => {
  },
  changeLanguage: () => {
  },
  currentLanguage: '',
};

const mapStateToProps = state => ({
  currentLanguage: state.ui.language,
});

const mapDispatchToProps = dispatch => ({
  authCheckState: () => dispatch(authCheckStateRedux()),
  changeLanguage: lng => dispatch(changeLanguageRedux(lng)),
});


export default connect(mapStateToProps, mapDispatchToProps)(RouterWrapper);
