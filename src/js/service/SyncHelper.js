import { Set } from 'immutable';
import { HOST } from '../configs';
import store from '../index';
import NotificationHelper from './NotificationHelper';

const SyncHelper = {};

SyncHelper.saveSelectionSnapshot = (participantsList) => {
  const selectionSnapshot = Set(participantsList.filter((participant) => {
    return participant.get('isSelected');
  }).map((participant) => {
    return participant.get('id');
  }));

  return fetch(`${HOST}/save`, {
    method: 'POST',
    headers: new Headers({ 'Content-Type': 'application/json' }),
    body: JSON.stringify({
      token: store.getState().get('auth').get('token'),
      state: JSON.stringify(selectionSnapshot),
    }),
  }).then((response) => {
    if (response.ok) {
      return response.json();
    }
    throw new Error('POST /save connection lost');
  }).catch(() => {
    NotificationHelper.showModalWithMessage('Connection lost. Please reload this page.');
  });
};

SyncHelper.fetchSelectionSnapshot = () => {
  return fetch(`${HOST}/load`, {
    headers: new Headers({ 'Content-Type': 'application/json' }),
    method: 'POST',
    body: JSON.stringify({
      token: store.getState().get('auth').get('token'),
    }),
  }).then((response) => {
    if (response.ok) {
      return response.json(); // should be resolved to immutable.js Set
    }
    throw new Error('POST /load connection lost');
  }).then((json) => {
    const state = JSON.parse(json.state);
    if (!Set.isSet(state)) {
      return Set([]);
    } else {
      return state;
    }
  }).catch(() => {
    NotificationHelper.showModalWithMessage('Connection lost. Please reload this page.');
  });
};

export default SyncHelper;
