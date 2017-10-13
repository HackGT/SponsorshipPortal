import { Set } from 'immutable';
import { HOST } from '../configs';
import store from '../index';
import NotificationHelper from './NotificationHelper';

const SyncHelper = {};

SyncHelper.saveSelectionSnapshot = () => {
  const participantsList = store.getState().get('participants');
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
      // Now that it has been saved, remove the 'beforeunload' listener that prevents user from exiting
      window.removeEventListener('beforeunload', SyncHelper.preventExitHelper);
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
    if (!json.state) {
      return Set([]);
    } else {
      return Set(JSON.parse(json.state));
    }
  }).catch(() => {
    NotificationHelper.showModalWithMessage('Connection lost. Please reload this page.');
  });
};

// Called when sync is needed. Update the UI and prevent users from accidentally exiting.
SyncHelper.requestSync = () => {
  // This function should remain idempotent
  if (!window.needSync) {
    // Prevent users from accidentally leaving the page
    window.addEventListener('beforeunload', SyncHelper.preventExitHelper);
    window.needSync = true;
    NotificationHelper.updateSyncStatus('Saving...');
  }
};

// Helper function used with window.addEventListener('beforeunload', ...) to prevent user from accidentally exiting the page
SyncHelper.preventExitHelper = (e) => {
  const dialogText = 'You are leaving the Sponsorship Portal. Are you sure?';
  e.returnValue = dialogText;
  return dialogText;
};

export default SyncHelper;
