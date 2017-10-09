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
