import 'semantic-ui-css/semantic.min.css';
import 'whatwg-fetch';
import { Map } from 'immutable';
import 'particles.js';
import particlesOption from '../assets/particlesOption';
import { configureHistory, configureStore } from './configureStoreAndHistory';
import '../css/index.css';

import render from './Routes';

const initialState = Map();
const store = configureStore(initialState);
const reactRouterReduxHistory = configureHistory(store);

render(store, reactRouterReduxHistory);

export default store;

// Particle Js Decorations
window.particlesJS('particle', particlesOption);
