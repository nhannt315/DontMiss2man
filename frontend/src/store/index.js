import {applyMiddleware, combineReducers, compose, createStore} from 'redux';
import {routerReducer, routerMiddleware} from 'react-router-redux';
import createSagaMiddleware from 'redux-saga';
import {watchBuilding, watchRoom, watchAuth} from './sagas';
import reducers from './reducers';

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
    routing: routerReducer,
  });

  const store = createStore(
    reducer,
    composeEnhancers(applyMiddleware(sagaMiddleware, routerMiddleware(history))),
  );
  store.sagaTask = sagaMiddleware.run(watchBuilding);
  store.sagaTask = sagaMiddleware.run(watchRoom);
  store.sagaTask = sagaMiddleware.run(watchAuth);
  return store;
};

export default configureStore;
