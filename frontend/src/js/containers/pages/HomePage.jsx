import React from 'react';
import { connect } from 'react-redux';
import { Container, Step } from 'semantic-ui-react';
import { TITLE } from '../../configs';

import ParticipantProfilePage from './ParticipantProfilePage';
import ExportPage from './ExportPage';

class HomePage extends React.Component {
  render() {
    // For sementics ui Step component
    const steps = [
      { icon: 'search', title: 'Search', description: 'Find the Talent you need' },
      { icon: 'plus', title: 'Select', description: 'Add Candidates to Your Selection' },
      { icon: 'cloud download', title: 'Export', description: 'Download Selected Data and Contacts' },
    ];

    return (
      <div>
        <Container textAlign="center" style={{ marginBottom: '40px' }}>
          <h1>{TITLE}</h1>
          {/* <h3>Access, Manage, Recruit and Contact Talents</h3> */}
          <Step.Group size="small" stackable="tablet" items={steps} />
        </Container>
        <ParticipantProfilePage />
        <ExportPage />
      </div>
    );
  }
}

export default connect()(HomePage);
