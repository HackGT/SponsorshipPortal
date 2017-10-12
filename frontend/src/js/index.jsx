import 'semantic-ui-css/semantic.min.css';
import 'whatwg-fetch';
import { Map } from 'immutable';
import 'particles.js';
// import particlesOption from '../assets/particlesOption';
import { configureHistory, configureStore } from './configureStoreAndHistory';
import '../css/index.css';
import { logInWithToken } from './actions/auth';

import render from './Routes';

// initializae store and router
const initialState = Map();
const store = configureStore(initialState);
const reactRouterReduxHistory = configureHistory(store);

// Starts rendering
render(store, reactRouterReduxHistory);

/**
 * Expose store globally
 */
export default store;

// Particle Js Decorations (Disabled)
// window.particlesJS('particle', particlesOption);

// Log the user in when token is present in localStorage
if (window.localStorage.getItem('token')) {
  store.dispatch(logInWithToken(window.localStorage.getItem('token')));
}
