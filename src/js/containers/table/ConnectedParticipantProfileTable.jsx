import { connect } from 'react-redux';
import ParticipantProfileTable from '../../components/table/ParticipantProfileTable';
import * as actions from '../../actions/participants';

function mapStateToProps(state) {
  return {
    participants: state.get('participants'),
  };
}

function mapDispatchToProps(dispatch) {
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
}

export default connect(mapStateToProps, mapDispatchToProps)(ParticipantProfileTable);
