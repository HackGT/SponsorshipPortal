import { List, Map, Set } from 'immutable';
import * as ACTION_TYPE from '../actionTypes';

// Example list of participants
const initialState = List([
  Map({
    id: 'loading',
    name: 'Loading',
    school: '...',
    email: '...',
    isSelected: false,
    isDisplaying: true,
    isSearched: false,
  }),
]);

export default function reducer(state = initialState, action) {
  switch (action.type) {
    case ACTION_TYPE.LOAD_PARTICIPANTS: {
      return action.payload;
    }
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
    case ACTION_TYPE.SEARCH_PARTICIPANTS: {
      const searchedSet = Set(action.payload);
      return state.map((participant) => {
        if (searchedSet.has(participant.get('id'))) {
          return participant.set('isDisplaying', true).set('isSearched', true);
        }
        return participant.set('isDisplaying', false).set('isSearched', false);
      });
    }
    default:
      return state;
  }
}
