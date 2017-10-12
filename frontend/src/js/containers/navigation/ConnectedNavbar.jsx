import { connect } from 'react-redux';
import { push } from 'react-router-redux';
import { logOut } from '../../actions/auth';
import Navbar from '../../components/navigation/Navbar';

function mapStateToProps(state, ownProps) {
  return {
    location: ownProps.location,
    ui: state.get('ui'),
    auth: state.get('auth'),
  };
}

function mapDispatchToProps(dispatch) {
  return {
    push: (uri) => {
      dispatch(push(uri));
    },
    logOut: () => {
      dispatch(logOut());
    },
  };
}

export default connect(mapStateToProps, mapDispatchToProps)(Navbar);
