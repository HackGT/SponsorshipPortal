import 'semantic-ui-css/semantic.min.css';

import React from 'react';
import ReactDOM from 'react-dom';

import { Map } from 'immutable';
import { Provider } from 'react-redux';
import { Router, Route } from 'react-router';

import { configureHistory, configureStore } from './configureStoreAndHistory';

import HomePage from './containers/pages/HomePage';
import ParticipantProfilePage from './containers/pages/ParticipantProfilePage';

const initialState = Map();
const store = configureStore(initialState);
const reactRouterReduxHistory = configureHistory(store);

ReactDOM.render((
  <Provider store={store}>
    <Router history={reactRouterReduxHistory}>
      <Route path="/" component={HomePage} />
      <Route path="/participant-profile" component={ParticipantProfilePage} />
    </Router>
  </Provider>
), document.getElementById('app'));
