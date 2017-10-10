/**
 * action creators of participants
 * 
 * All parameters using Immutable JS except loadParticipantsListFromArray
 */
import { List, Map } from 'immutable';
import * as ACTION_TYPES from '../actionTypes';
import { HOST } from '../configs';
import { loaderOff } from './ui';
import store from '../index';
import SyncHelper from '../service/SyncHelper';
import NotificationHelper from '../service/NotificationHelper';


// Fetch participants from REST API
export function loadParticipants() {
  return (dispatch) => {
    fetch(`${HOST}/participants`, {
      token: store.getState().get('auth').get('token'),
      method: 'POST',
    }).then((response) => {
      if (response.ok) {
        return response.json();
      }
      throw new Error('Lost server connection');
    }).then((json) => {
      SyncHelper.fetchSelectionSnapshot().then((selectedIdSet) => {
        // Syncing selection state from server and merge into participantsList
        const participantsList = List(Object.entries(json).map((entry) => {
          return Map({
            id: entry[0],
            name: entry[1].name,
            email: entry[1].email,
            resumeId: entry[1].resumeId,
            isSelected: selectedIdSet.has(entry[0]),
            isSearched: false,
            isDisplaying: true,
          });
        }));
        dispatch({
          type: ACTION_TYPES.LOAD_PARTICIPANTS,
          payload: participantsList,
        });
        dispatch(loaderOff()); // end of the login initialization workflow, turn off the loader
      });
    }).catch(() => {
      NotificationHelper.showModalWithMessage('Connection lost. Please reload this page.');
    });
  };
}

export function selectParticipant(participant) {
  return (dispatch) => {
    dispatch({
      type: ACTION_TYPES.SELECT_PARTICIPANT,
      payload: participant,
    });
  };
}

export function unSelectParticipant(participant) {
  return (dispatch) => {
    dispatch({
      type: ACTION_TYPES.UNSELECT_PARTICIPANT,
      payload: participant,
    });
  };
}

export function showParticipant(participant) {
  return (dispatch) => {
    dispatch({
      type: ACTION_TYPES.SHOW_PARTICIPANT,
      payload: participant,
    });
  };
}

export function hideParticipant(participant) {
  return (dispatch) => {
    dispatch({
      type: ACTION_TYPES.HIDE_PARTICIPANT,
      payload: participant,
    });
  };
}
