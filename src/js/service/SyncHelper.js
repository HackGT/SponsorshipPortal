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
    token: store.getState().get('auth').get('token'),
    method: 'POST',
    state: JSON.stringify(selectionSnapshot),
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
    token: store.getState().get('auth').get('token'),
    method: 'POST',
  }).then((response) => {
    if (response.ok) {
      return response.json(); // should be immutable.js Set
    }
    throw new Error('POST /load connection lost');
  }).catch(() => {
    NotificationHelper.showModalWithMessage('Connection lost. Please reload this page.');
  });
};

export default SyncHelper;
