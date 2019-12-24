import React from 'react';
import {Provider} from 'react-redux';
import {createBrowserHistory} from 'history';
import './styles/base.scss';
import AppRoutes from './config/routes';
import configureStore from './store';
import configureAxios from './config/axios';

const App = () => {
  const history = createBrowserHistory();
  const store = configureStore(history);
  configureAxios(store);
  return (
    <Provider store={store}>
      <AppRoutes history={history} />
    </Provider>
  );
};

export default App;
