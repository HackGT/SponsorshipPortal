import 'semantic-ui-css/semantic.min.css';
import { Map } from 'immutable';
import { configureHistory, configureStore } from './configureStoreAndHistory';

import render from './Routes';

const initialState = Map();
const store = configureStore(initialState);
const reactRouterReduxHistory = configureHistory(store);

render(store, reactRouterReduxHistory);
