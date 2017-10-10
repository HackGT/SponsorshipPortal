import store from '../index';
import * as ACTION_TYPES from '../actionTypes';
import { updateSyncStatus } from '../actions/ui';

const NotificationHelper = {};

NotificationHelper.showModalWithMessage = (message) => {
  store.dispatch({
    type: ACTION_TYPES.MODAL_ON,
    payload: {
      message,
    },
  });
};

NotificationHelper.updateSyncStatus = (message) => {
  store.dispatch(updateSyncStatus(message));
};

export default NotificationHelper;
