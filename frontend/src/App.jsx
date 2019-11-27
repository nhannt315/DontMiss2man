import React from 'react';
import {Provider} from 'react-redux';
import {createBrowserHistory} from 'history';

import './styles/base.scss';
import Layout from './components/Layout';
import AppRoutes from './config/routes';
import configureStore from './store';
import configureAxios from './config/axios';

const App = () => {
  const history = createBrowserHistory();
  const store = configureStore(history);
  configureAxios();
  return (
    <Provider store={store}>
      <Layout>
        <AppRoutes history={history} />
      </Layout>
    </Provider>
  );
};

export default App;
