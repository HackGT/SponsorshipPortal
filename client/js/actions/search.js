import { Set } from 'immutable';
import * as ACTION_TYPES from '../actionTypes';
import { tableLoaderOn, tableLoaderOff } from './ui';
import { HOST } from '../configs';
import store from '../index';
import NotificationHelper from '../service/NotificationHelper';

export function searchByKeyword(keyword) {
  return (dispatch) => {
    dispatch(tableLoaderOn());

    return fetch(`${HOST}/search`, {
      method: 'POST',
      headers: new Headers({ 'Content-Type': 'application/json' }),
      body: JSON.stringify({
        query: keyword,
        token: store.getState().get('auth').get('token'),
      }),
    }).then((response) => {
      if (response.ok) {
        return response.json(); // should be immutable.js Set
      }
      throw new Error('POST /search connection lost');
    }).then((json) => {
      if (!json.results.hits) {
        throw new Error('Invalid search response');
      }
      const searchedIdSet = Set(json.results.hits.map(hit => hit.id));
      // Update participants state
      dispatch({
        type: ACTION_TYPES.SEARCH_PARTICIPANTS,
        payload: searchedIdSet,
      });
      // Change the Currently Cached Keyword for display
      dispatch(updateCurrentKeyword(keyword));

      // Finish loading
      dispatch(tableLoaderOff());
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
