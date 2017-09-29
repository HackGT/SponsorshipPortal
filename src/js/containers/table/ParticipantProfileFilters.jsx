import React from 'react';
import { connect } from 'react-redux';
import { Button } from 'semantic-ui-react';
import * as actions from '../../actions/participants';

class ParticipantProfileFilters extends React.Component {
  render() {
    const participants = this.props.participants;
    const showParticipant = this.props.showParticipant;
    const hideParticipant = this.props.hideParticipant;
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
      </Button.Group>
    );
  }
}

export default connect((state) => {
  return {
    participants: state.get('participants'),
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
