import React from 'react';
import { connect } from 'react-redux';
import { Container, Step } from 'semantic-ui-react';

class HomePage extends React.Component {
  render() {
    // For sementics ui Step component
    const steps = [
      { icon: 'search', title: 'Search', description: 'Find the Talent you need' },
      { icon: 'plus', title: 'Select', description: 'Add Candidates to Your Selection' },
      { icon: 'cloud download', title: 'Export', description: 'Download Selected Data and Contacts' },
    ];

    return (
      <Container textAlign="center">
        <h1>Welcome to HackGT Sponsorship Portal</h1>
        <h3>Access, Manage, Recruit and Contact Talents</h3>
        <Step.Group size="small" items={steps} />
      </Container>
    );
  }
}

export default connect()(HomePage);
