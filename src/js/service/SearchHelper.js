const SearchHelper = {};

/**
 * Return an immutable copy of the new participantsList after merging searchQuery data
 * 
 * @param processedSearchQuery an Immutable.js Set processed from the search result JSON, containing ids of matching participants
 * @param participantsList an Immutable.js List from the redux state.participants
 * @return An immutable copy of the new participantsList to be used for state.participants
 */
SearchHelper.mergeSearchToParticipantsList = (processedSearchResult, participantsList) => {
  return participantsList.map((participant) => {
    if (processedSearchResult.includes(participant.get('id'))) {
      return participant.set('isSearched', true);
    }
    return participant.set('isSearched', false);
  });
};

// SearchHelper.preprocessSearchResult = (rawSearchResult) => {

// };

// SearchHelper.requestSearch = (keyword) => {

// };

export default SearchHelper;
