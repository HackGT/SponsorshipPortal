import * as ACTION_TYPES from '../actionTypes';
import { loaderOn, loaderOff } from './ui';
import { loadParticipants } from './participants';
import { HOST, IS_DEV_ENV } from '../configs';
import NotificationHelper from '../service/NotificationHelper';
import SyncHelper from '../service/SyncHelper';

export function logIn(username, password) {
  if (IS_DEV_ENV) {
    return logInDev();
  }
  // Request Server Auth login
  return (dispatch) => {
    dispatch(loaderOn());
    return fetch(`${HOST}/login`, {
      method: 'POST',
      headers: new Headers({ 'Content-Type': 'application/json' }),
      body: JSON.stringify({
        username,
        password,
      }),
    }).then((response) => {
      if (response.ok) {
        return response.json();
      }
      throw new Error('Login Failed');
    }).then((json) => {
      if (!json.token) {
        throw new Error('Login Failed due to invalid credentials');
      }
      // Update redux state with token, which is needed for all subsequent API requests
      dispatch({
        type: ACTION_TYPES.LOG_IN,
        payload: {
          // username,
          token: json.token,
        },
      });
      // Save token to localStorage
      window.localStorage.setItem('token', json.token);
      // Fetch participants and sync selection state
      dispatch(loadParticipants());
      // Initialize Auto Syncing. See SyncHelper.requestSync for more info.
      window.needSync = false;
      window.setInterval(() => {
        if (window.needSync) {
          SyncHelper.saveSelectionSnapshot();
          NotificationHelper.updateSyncStatus('Progress Saved');
          window.needSync = false;
        }
      }, 3000);
    }).catch(() => {
      dispatch(loaderOff()); // let user try another credential, prevent the loader/dimmer from not shutting down
      NotificationHelper.showModalWithMessage('Login Failure: Please check your credentials');
    });
  };
}

export function logInWithToken(token) {
  return (dispatch) => {
    dispatch(loaderOn());
    dispatch({
      type: ACTION_TYPES.LOG_IN,
      payload: {
        token,
      },
    });

    // Fetch participants and sync selection state; If the token is expired, this will also log it out
    dispatch(loadParticipants());
    // Initialize Auto Syncing. See SyncHelper.requestSync for more info.
    window.needSync = false;
    window.setInterval(() => {
      if (window.needSync) {
        SyncHelper.saveSelectionSnapshot();
        NotificationHelper.updateSyncStatus('Progress Saved');
      }
    }, 3000);
    dispatch(loaderOff());
  };
}

// logIn used only for frontend developement and experiments
function logInDev() {
  // mock token response
  const json = { token: 'test_token' };

  return (dispatch) => {
    dispatch(loaderOn());
    dispatch({
      type: ACTION_TYPES.LOG_IN,
      payload: {
        token: json.token,
      },
    });
    // Save token to localStorage
    window.localStorage.setItem('token', json.token);
    dispatch(loadParticipants());
    dispatch(loaderOff());
  };
}

export function logOut() {
  return (dispatch) => {
    dispatch({
      type: ACTION_TYPES.LOG_OUT,
    });
    window.localStorage.clear(); // clean expired token, if any
  };
}
