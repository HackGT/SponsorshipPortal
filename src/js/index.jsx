import 'semantic-ui-css/semantic.min.css';
import 'whatwg-fetch';
import { Map } from 'immutable';
import 'particles.js';
import particlesOption from '../assets/particlesOption';
import { configureHistory, configureStore } from './configureStoreAndHistory';
import '../css/index.css';

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

// Particle Js Decorations
window.particlesJS('particle', particlesOption);

// Prevent users from accidentally leaving the page
window.onbeforeunload = (e) => {
  const dialogText = 'You are leaving the Sponsorship Portal. Are you sure?';
  e.returnValue = dialogText;
  return dialogText;
};
