import store from '../index';
import * as ACTION_TYPES from '../actionTypes';

const NotificationHelper = {};

NotificationHelper.showModalWithMessage = (message) => {
  store.dispatch({
    type: ACTION_TYPES.MODAL_ON,
    payload: {
      message,
    },
  });
};

export default NotificationHelper;
