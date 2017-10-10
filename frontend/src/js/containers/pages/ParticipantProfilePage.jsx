import React from 'react';
import { connect } from 'react-redux';
import { Container } from 'semantic-ui-react';

import ConnectedParticipantProfileTable from '../table/ConnectedParticipantProfileTable';

class ParticipantProfilePage extends React.Component {
  render() {
    return (
      <Container style={{ marginBottom: '40px' }}>
        <ConnectedParticipantProfileTable />
      </Container>
    );
  }
}

export default connect()(ParticipantProfilePage);