import { List, Map } from 'immutable';
import * as ACTION_TYPE from '../actionTypes';

// Example list of participants
const initialState = List([
  Map({
    id: 1,
    name: 'John Doe',
    school: 'Georgia Tech',
    email: 'john@example.com',
    isSelected: false,
    isDisplaying: true,
  }),
  Map({
    id: 2,
    name: 'John Doe 2',
    school: 'Georgia Tech',
    email: 'john2@example.com',
    isSelected: true,
    isDisplaying: true,
  }),
  Map({
    id: 3,
    name: 'John Doe 3',
    school: 'Georgia Tech',
    email: 'john3@example.com',
    isSelected: false,
    isDisplaying: true,
  }),
]);

export default function reducer(state = initialState, action) {
  switch (action.type) {
    case ACTION_TYPE.SELECT_PARTICIPANT: {
      return state.update(state.findIndex((participant) => {
        return participant.get('id') === action.payload.get('id');
      }), (participant) => {
        return participant.set('isSelected', true);
      });
    }
    case ACTION_TYPE.UNSELECT_PARTICIPANT: {
      return state.update(state.findIndex((participant) => {
        return participant.get('id') === action.payload.get('id');
      }), (participant) => {
        return participant.set('isSelected', false);
      });
    }
    case ACTION_TYPE.SHOW_PARTICIPANT: {
      return state.update(state.findIndex((participant) => {
        return participant.get('id') === action.payload.get('id');
      }), (participant) => {
        return participant.set('isDisplaying', true);
      });
    }
    case ACTION_TYPE.HIDE_PARTICIPANT: {
      return state.update(state.findIndex((participant) => {
        return participant.get('id') === action.payload.get('id');
      }), (participant) => {
        return participant.set('isDisplaying', false);
      });
    }
    default:
      return state;
  }
}
