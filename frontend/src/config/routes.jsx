import React from 'react';
import {Switch, Router, Route} from 'react-router';
import PropTypes from 'prop-types';
import Layout from '../components/Layout';

import * as Containers from '../containers';

const buildRoutes = history => {
  return (
    <Router history={history}>
      <Layout>
        <Switch>
          <Route exact path="/" component={Containers.HomePage} />
          <Route exact path="/chintai/:id" component={Containers.RoomDetailPage} />
        </Switch>
      </Layout>
    </Router>
  );
};

const RouterWrapper = ({history}) => {
  return <div className="router-wrapper">{buildRoutes(history)}</div>;
};

RouterWrapper.propTypes = {
  history: PropTypes.object,
};

RouterWrapper.defaultProps = {
  history: {},
};

export default RouterWrapper;
