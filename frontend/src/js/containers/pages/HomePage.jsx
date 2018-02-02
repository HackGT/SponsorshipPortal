import React from 'react';
import { connect } from 'react-redux';

import ParticipantProfilePage from './ParticipantProfilePage';

class HomePage extends React.Component {
  render() {
    return (
      <div>
        <ParticipantProfilePage />
      </div>
    );
  }
}

export default connect()(HomePage);
