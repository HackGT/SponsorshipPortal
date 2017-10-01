/**
 * action creators of search
 * 
 * All parameters using Immutable JS
 */
// import { List, Map } from 'immutable';
import * as ACTION_TYPES from '../actionTypes';


export function searchByKeyword(keyword) {
  return (dispatch) => {
    if (keyword === 'John Doe 2') {
      dispatch({
        type: ACTION_TYPES.SEARCH_PARTICIPANTS,
      });
    }
    dispatch({
      type: ACTION_TYPES.SEARCH_BY_KEYWORD,
      payload: keyword,
    });
  };
}

export function cacheKeyword() {

}
