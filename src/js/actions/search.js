/**
 * action creators of search
 * 
 * All parameters using Immutable JS
 */
// import { List, Map } from 'immutable';
import * as ACTION_TYPES from '../actionTypes';


export function searchByKeyword(keyword) {
  return (dispatch) => {
    // Change the Currently Cached Keyword
    dispatch({
      type: ACTION_TYPES.SEARCH_BY_KEYWORD,
      payload: keyword,
    });
    // Show Loader
    dispatch({
      type: ACTION_TYPES.LOADER_ON,
    });
    // Send Search Request
    dispatch({
      type: ACTION_TYPES.SEARCH_PARTICIPANTS,
      payload: ['1', '2'],
    });
  };
}

export function cacheKeyword() {

}
