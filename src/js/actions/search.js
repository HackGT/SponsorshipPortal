import { Set } from 'immutable';
import * as ACTION_TYPES from '../actionTypes';
import { loaderOn, loaderOff } from './ui';
import { HOST } from '../configs';
import store from '../index';

export function searchByKeyword(keyword) {
  return (dispatch) => {
    dispatch(loaderOn());

    return fetch(`${HOST}/search`, {
      token: store.getState().get('auth').get('token'),
      method: 'POST',
      query: keyword,
    }).then((response) => {
      if (response.ok) {
        return response.json(); // should be immutable.js Set
      }
      throw new Error('POST /search connection lost');
    }).then((json) => {
      if (!json.hits) {
        throw new Error('Invalid search response');
      }
      const searchedIdSet = Set(json.hits.map(hit => hit.id));
      // Update participants state
      dispatch({
        type: ACTION_TYPES.SEARCH_PARTICIPANTS,
        payload: searchedIdSet,
      });
      // Change the Currently Cached Keyword for display
      updateCurrentKeyword(keyword);

      // Finish loading
      dispatch(loaderOff());
    }).catch(() => {
      NotificationHelper.showModalWithMessage('Connection lost. Please reload this page.');
    });
  };
}

// Change the Currently Cached Keyword for display
export function updateCurrentKeyword(keyword) {
  return (dispatch) => {
    dispatch({
      type: ACTION_TYPES.SEARCH_BY_KEYWORD,
      payload: keyword,
    });
  };
}
