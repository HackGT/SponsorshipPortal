import { Map } from 'immutable';
import * as ACTION_TYPE from '../actionTypes';

const initialState = Map({
  keyword: '',
});

export default function reducer(state = initialState, action) {
  switch (action.type) {
    case ACTION_TYPE.SEARCH_BY_KEYWORD: {
      return state.set('keyword', action.payload);
    }
    default:
      return state;
  }
}
