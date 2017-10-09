import * as ACTION_TYPES from '../actionTypes';
import { loaderOn, loaderOff } from './ui';
import { HOST } from '../configs';
import NotificationHelper from '../service/NotificationHelper';

export function logIn(username, password) { // eslint-disable-line import/prefer-default-export
  // Request Server Auth login
  return (dispatch) => {
    dispatch(loaderOn());
    return fetch(`${HOST}/login`, {
      method: 'POST',
      body: {
        username,
        password,
      },
    }).then((response) => {
      dispatch(loaderOff());
      if (response.ok) {
        return response.json();
      }
      throw new Error('Login Failed');
    }).then((json) => {
      if (!json.token) {
        throw new Error('Login Failed due to invalid credentials');
      }
      // Update redux state with token, which is needed for all subsequent API requests
      return dispatch({
        type: ACTION_TYPES.LOG_IN,
        payload: {
          username,
          token: json.token,
        },
      });
    }).catch(() => {
      dispatch(loaderOff()); // let user try another credential, prevent the loader/dimmer from not shutting down
      NotificationHelper.showModalWithMessage('Login Failure: Please check your credentials');
    });
  };
}
