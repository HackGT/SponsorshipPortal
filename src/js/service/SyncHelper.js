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
    headers: { 'Content-Type': 'application/json' },
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
    headers: { 'Content-Type': 'application/json' },
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
    if (!json.state) {
      throw new Error('Invalid snapshot');
    }
    if (json.state === 'none') {
      return Promise.resolve(Set([]));
    } else {
      return Promise.resolve(JSON.parse(json.state));
    }
  }).catch(() => {
    NotificationHelper.showModalWithMessage('Connection lost. Please reload this page.');
  });
};

export default SyncHelper;
