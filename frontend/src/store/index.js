import { applyMiddleware, combineReducers, compose, createStore } from 'redux';
import {routerReducer, routerMiddleware} from 'react-router-redux';
import reducers from './reducers';
import createSagaMiddleware from 'redux-saga';

const composeEnhancers =
  typeof window === 'object' &&
  window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ ?
    window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__({
      // Specify extension’s options like name, actionsBlacklist, actionsCreators, serialize...
    }) : compose;

const sagaMiddleware = createSagaMiddleware();

const configureStore = history => {
  const reducer = combineReducers({
    ...reducers,
    routing: routerReducer
  });

  return createStore(
    reducer,
    composeEnhancers(applyMiddleware(sagaMiddleware, routerMiddleware(history)))
  );
};

export default configureStore;
