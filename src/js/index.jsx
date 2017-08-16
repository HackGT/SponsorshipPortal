import 'semantic-ui-css/semantic.min.css';

import React from 'react';
import ReactDOM from 'react-dom';

import { Map } from 'immutable';
import { Provider } from 'react-redux';
import { Router, Route } from 'react-router';

import { configureHistory, configureStore } from './configureStoreAndHistory';

import Homepage from './Homepage/containers/Homepage';

const initialState = Map();
const store = configureStore(initialState);
const reactRouterReduxHistory = configureHistory(store);

ReactDOM.render((
  <Provider store={store}>
    <Router history={reactRouterReduxHistory}>
      <Route path="/" component={Homepage} />
    </Router>
  </Provider>
), document.getElementById('app'));
