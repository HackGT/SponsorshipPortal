import { combineReducers } from 'redux-immutable';

import { routerReducer } from 'react-router-redux';
import ExampleReducer from './ExampleReducer';
import participants from './participants';

export default combineReducers({
  routing: routerReducer,
  participants,
  example: ExampleReducer,
});
