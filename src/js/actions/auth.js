import * as ACTION_TYPES from '../actionTypes';

export function logIn(username, password) { // eslint-disable-line import/prefer-default-export
  // Request Server Auth login
  return {
    type: ACTION_TYPES.LOG_IN,
  };
}
