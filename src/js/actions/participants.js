import * as ACTION_TYPES from '../actionTypes';

export function selectParticipant(participant) {
  return (dispatch) => {
    dispatch({
      type: ACTION_TYPES.SELECT_PARTICIPANT,
      payload: participant,
    });
  };
}

export function unSelectParticipant(participant) {
  return (dispatch) => {
    dispatch({
      type: ACTION_TYPES.UNSELECT_PARTICIPANT,
      payload: participant,
    });
  };
}

export function showParticipant(participant) {
  return (dispatch) => {
    dispatch({
      type: ACTION_TYPES.SHOW_PARTICIPANT,
      payload: participant,
    });
  };
}

export function hideParticipant(participant) {
  return (dispatch) => {
    dispatch({
      type: ACTION_TYPES.HIDE_PARTICIPANT,
      payload: participant,
    });
  };
}
