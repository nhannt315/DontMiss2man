import React from 'react';
import {Provider} from 'react-redux';
import ReactGA from 'react-ga';
import {createBrowserHistory} from 'history';
import './styles/base.scss';
import AppRoutes from './config/routes';
import configureStore from './store';
import configureAxios from './config/axios';
const history = createBrowserHistory();
history.listen(location => {
  ReactGA.set({ page: location.pathname }); // Update the user's current page
  ReactGA.pageview(location.pathname); // Record a pageview for the given page
});

const App = () => {
  
  const store = configureStore(history);
  configureAxios(store);

  return (
    <Provider store={store}>
      <AppRoutes history={history} />
    </Provider>
  );
};

export default App;
