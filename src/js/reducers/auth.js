import { Map } from 'immutable';
import * as ACTION_TYPE from '../actionTypes';

const initialState = Map({
  active: false,
  username: null,
  token: null,
});

export default function reducer(state = initialState, action) {
  switch (action.type) {
    case ACTION_TYPE.LOG_IN: {
      return Map({
        active: true,
        username: action.payload.username,
        token: action.payload.token,
      });
    }
    default:
      return state;
  }
}
