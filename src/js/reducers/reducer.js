import { combineReducers } from 'redux-immutable';

import { routerReducer } from 'react-router-redux';
import ExampleReducer from './ExampleReducer';
import participants from './participants';
import ui from './ui';
import search from './search';

export default combineReducers({
  routing: routerReducer,
  participants,
  ui,
  search,
  example: ExampleReducer,
});
