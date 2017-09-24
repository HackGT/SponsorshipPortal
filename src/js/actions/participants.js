/**
 * action creators of participants
 * 
 * All parameters using Immutable JS except loadParticipantsListFromArray
 */
import { List, Map } from 'immutable';
import * as ACTION_TYPES from '../actionTypes';


export function loadParticipantsListFromArray(participantsArray) {
  return loadParticipantsList(List(participantsArray).map((participant) => {
    return Map(participant);
  }));
}

export function loadParticipantsList(participantsList) {
  if (!List.isList(participantsList)) {
    throw new Error('Action Creators Require Immutable.JS Parameters');
  }
  return (dispatch) => {
    dispatch({
      type: ACTION_TYPES.LOAD_PARTICIPANTS,
      payload: participantsList,
    });
  };
}

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
