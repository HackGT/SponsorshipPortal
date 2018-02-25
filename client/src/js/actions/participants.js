/**
 * action creators of participants
 * 
 * All parameters using Immutable JS except loadParticipantsListFromArray
 */
import { List, Map } from 'immutable';
import * as ACTION_TYPES from '../actionTypes';
import { HOST, IS_DEV_ENV } from '../configs';
import { loaderOff } from './ui';
import { logOut } from './auth';
import store from '../index';
import SyncHelper from '../service/SyncHelper';
import NotificationHelper from '../service/NotificationHelper';


// Fetch participants from REST API
export function loadParticipants() {
  if (IS_DEV_ENV) {
    return loadParticipantsDev();
  }
  return (dispatch) => {
    fetch(`${HOST}/participants`, {
      method: 'POST',
      headers: new Headers({ 'Content-Type': 'application/json' }),
      body: JSON.stringify({
        token: store.getState().get('auth').get('token'),
      }),
    }).then((response) => {
      if (response.ok) {
        return response.json();
      }
      throw new Error('Lost server connection');
    }).then((json) => {
      if (json.error) {
        // token expired, log the user out
        dispatch(logOut());
        return;
      }
      // If token correct
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

function loadParticipantsDev() {
  // mock response
  const json = {
    1640627800689991751: {
      email: 'anish.visaria@gatech.edu',
      name: 'Anish Visaria',
      resumeId: 'dfdsfsdfds',
    },
    2904581336711582906: {
      email: 'robert@gatech.edu',
      name: 'Robert Li',
      resumeId: 'dfdsfsdfds',
    },
  };

  return (dispatch) => {
    const participantsList = List(Object.entries(json).map((entry) => {
      return Map({
        id: entry[0],
        name: entry[1].name,
        email: entry[1].email,
        resumeId: entry[1].resumeId,
        isSelected: false, // This is stateless dummy data
        isSearched: false,
        isDisplaying: true,
      });
    }));
    dispatch({
      type: ACTION_TYPES.LOAD_PARTICIPANTS,
      payload: participantsList,
    });
    dispatch(loaderOff());
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
