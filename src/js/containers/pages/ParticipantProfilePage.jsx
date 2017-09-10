import React from 'react';
import { connect } from 'react-redux';
import { Container } from 'semantic-ui-react';

import ConnectedParticipantProfileTable from '../table/ConnectedParticipantProfileTable';

import Navbar from '../../components/navigation/Navbar';

class ParticipantProfilePage extends React.Component {
  render() {
    return (
      <Container>
        <Navbar page="ParticipantProfilePage" />
        <ConnectedParticipantProfileTable />
      </Container>
    );
  }
}

export default connect()(ParticipantProfilePage);
