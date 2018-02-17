import * as ACTION_TYPES from '../actionTypes';

export default function changeMessage() {
  return (dispatch, getState) => {
    dispatch({
      type: ACTION_TYPES.CHANGING_MESSAGE,
    });
    const currentMessage = getState().getIn(['example', 'message']);
    let newMessage;
    switch (currentMessage.toLowerCase()) {
      case 'hello': {
        newMessage = 'world';
        break;
      }
      case 'world': {
        newMessage = 'hello';
        break;
      }
      default: {
        newMessage = 'hello';
      }
    }
    dispatch({
      type: ACTION_TYPES.CHANGED_MESSAGE,
      payload: newMessage,
    });
  };
}
