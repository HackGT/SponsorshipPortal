import React from 'react';
import { connect } from 'react-redux';
import { Button } from 'semantic-ui-react';
import * as actions from '../../actions/participants';

class ParticipantProfileFilters extends React.Component {
  render() {
    const participants = this.props.participants;
    const showParticipant = this.props.showParticipant;
    const hideParticipant = this.props.hideParticipant;
    const keyword = this.props.keyword;
    const searchedCount = participants.count(participant => participant.get('isSearched'));
    let searchFilterButton = (keyword === '') ? false : (
      <Button
        onClick={() => {
          // show all the participants currently matched by the search
          participants.forEach((participant) => {
            if (participant.get('isSearched')) {
              showParticipant(participant);
            } else {
              hideParticipant(participant);
            }
          });
        }}
      >
        Searching: {keyword} ({searchedCount})
      </Button>
    );

    if (keyword !== '' && searchedCount === 0) {
      searchFilterButton = (
        <Button color="olive">
          No Result Found
        </Button>
      );
    }

    return (
      <Button.Group>
        <Button
          onClick={() => {
            // show all the participants loaded
            participants.forEach((participant) => {
              showParticipant(participant);
            });
          }}
        >
          All ({participants.count()})
        </Button>
        <Button
          onClick={() => {
            // show all the participants selected
            participants.forEach((participant) => {
              if (participant.get('isSelected')) {
                showParticipant(participant);
              } else {
                hideParticipant(participant);
              }
            });
          }}
        >
          Selected ({participants.count(participant => participant.get('isSelected'))})
        </Button>
        {/* <Button
          onClick={() => {
            // fetch set of id of participants that the user has met
            participants.forEach((participant) => {
              if (participant.get('isSelected')) {
                showParticipant(participant);
              } else {
                hideParticipant(participant);
              }
            });
          }}
        >
          Met ({participants.count(participant => participant.get('isSelected'))})
        </Button> */}
        {searchFilterButton}
      </Button.Group>
    );
  }
}

export default connect((state) => {
  return {
    participants: state.get('participants'),
    keyword: state.get('search').get('keyword'),
  };
}, (dispatch) => {
  return {
    selectParticipant: (participant) => {
      dispatch(actions.selectParticipant(participant));
    },
    unSelectParticipant: (participant) => {
      dispatch(actions.unSelectParticipant(participant));
    },
    showParticipant: (participant) => {
      dispatch(actions.showParticipant(participant));
    },
    hideParticipant: (participant) => {
      dispatch(actions.hideParticipant(participant));
    },
  };
})(ParticipantProfileFilters);
