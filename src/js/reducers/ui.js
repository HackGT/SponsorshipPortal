import { Map } from 'immutable';
import * as ACTION_TYPE from '../actionTypes';

// Example list of participants
const initialState = Map({
  loader: false,
  modalActive: false,
  modalMessage: '',
  syncStatus: 'Please log in',
});

export default function reducer(state = initialState, action) {
  switch (action.type) {
    case ACTION_TYPE.LOADER_ON: {
      return state.set('loader', true);
    }
    case ACTION_TYPE.LOADER_OFF: {
      return state.set('loader', false);
    }
    case ACTION_TYPE.MODAL_ON: {
      return state.set('modalMessage', action.payload.message).set('modalActive', true);
    }
    case ACTION_TYPE.MODAL_OFF: {
      return state.set('modalActive', false);
    }
    case ACTION_TYPE.UI_UPDATE_SYNC_STATUS: {
      return state.set('syncStatus', action.payload.message);
    }
    default:
      return state;
  }
}
