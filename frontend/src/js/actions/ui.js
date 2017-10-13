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

export function tableLoaderOn() {
  return {
    type: ACTION_TYPES.TABLE_LOADER_ON,
  };
}

export function tableLoaderOff() {
  return {
    type: ACTION_TYPES.TABLE_LOADER_OFF,
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

export function showPDFAtURLInModal(url) {
  return (dispatch) => {
    dispatch({
      type: ACTION_TYPES.UI_UPDATE_PDF_URL,
      payload: {
        url,
      },
    });
    dispatch({
      type: ACTION_TYPES.UI_PDF_MODAL_ON,
    });
  };
}
