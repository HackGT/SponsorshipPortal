import { combineReducers } from 'redux-immutable';

import { routerReducer } from 'react-router-redux';
import ExampleReducer from './ExampleReducer';

export default combineReducers({
  routing: routerReducer,
  example: ExampleReducer,
});
