import { connect } from 'react-redux';
import { push } from 'react-router-redux';

import Navbar from '../../components/navigation/Navbar';

function mapStateToProps(state, ownProps) {
  return {
    location: ownProps.location,
  };
}

function mapDispatchToProps(dispatch) {
  return {
    push: (uri) => {
      dispatch(push(uri));
    },
  };
}

export default connect(mapStateToProps, mapDispatchToProps)(Navbar);
