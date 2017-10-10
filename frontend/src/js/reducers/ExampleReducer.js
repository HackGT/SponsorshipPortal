import { Map } from 'immutable';
import * as ACTION_TYPE from '../actionTypes';

const initialState = Map({
  message: 'hello',
  changingMessage: false,
});

export default function reducer(state = initialState, action) {
  switch (action.type) {
    case ACTION_TYPE.CHANGING_MESSAGE: {
      return state.merge({
        changingMessage: true,
      });
    }
    case ACTION_TYPE.CHANGED_MESSAGE: {
      return state.merge({
        changingMessage: false,
        message: action.payload,
      });
    }
    default:
      return state;
  }
}
