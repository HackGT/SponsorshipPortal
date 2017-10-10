import * as ACTION_TYPES from '../actionTypes';

export function loaderOn() {
  return {
    type: ACTION_TYPES.LOADER_ON,
  };
}

export function loaderOff() {
  return {
    type: ACTION_TYPES.LOADER_OFF,
  };
}

export function updateSyncStatus(message) {
  return {
    type: ACTION_TYPES.UI_UPDATE_SYNC_STATUS,
    payload: {
      message,
    },
  };
}
