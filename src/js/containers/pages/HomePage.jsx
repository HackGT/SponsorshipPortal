import React from 'react';
import { connect } from 'react-redux';
import { Container } from 'semantic-ui-react';

class HomePage extends React.Component {
  render() {
    return (
      <Container>
        <p>This page is under construction</p>
      </Container>
    );
  }
}

export default connect()(HomePage);
