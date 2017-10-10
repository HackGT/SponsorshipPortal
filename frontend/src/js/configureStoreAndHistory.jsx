import ReduxFreeze from 'redux-freeze';
import ReduxThunk from 'redux-thunk';
import logger from 'redux-logger';
import { browserHistory, hashHistory } from 'react-router';
import { createStore, applyMiddleware } from 'redux';
import { syncHistoryWithStore, routerMiddleware } from 'react-router-redux';

import rootReducer from './reducers/reducer';

export const history = process.env.DEVELOPMENT ? hashHistory : browserHistory;
let middleware = applyMiddleware(routerMiddleware(history), ReduxThunk, ReduxFreeze, logger);

if (!process.env.DEVELOPMENT) {
  middleware = applyMiddleware(routerMiddleware(history), ReduxThunk);
}

export function configureStore(initialState) {
  return createStore(rootReducer, initialState, middleware);
}

export function configureHistory(store) {
  return syncHistoryWithStore(
    history,
    store,
    { selectLocationState: (state => state.get('routing')) },
  );
}
