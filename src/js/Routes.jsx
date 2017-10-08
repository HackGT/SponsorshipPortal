import React from 'react';
import ReactDOM from 'react-dom';

import { Provider } from 'react-redux';
import { Router, Route, IndexRoute } from 'react-router';

import App from './App';
import HomePage from './containers/pages/HomePage';
import ParticipantProfilePage from './containers/pages/ParticipantProfilePage';
import ExportPage from './containers/pages/ExportPage';

const render = (store, reactRouterReduxHistory) => {
  ReactDOM.render((
    <Provider store={store}>
      <Router history={reactRouterReduxHistory}>
        <Route path="/" component={App}>
          <IndexRoute component={HomePage} />
          {/* <Route path="participant-profile" component={ParticipantProfilePage} /> */}
          <Route path="export" component={ExportPage} />
        </Route>
      </Router>
    </Provider>
  ), document.getElementById('app'));
};

export default render;

