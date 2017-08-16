import React from 'react';
import { connect } from 'react-redux';
import { Container, Header } from 'semantic-ui-react';

class Homepage extends React.Component {
  render() {
    return (
      <Container>
        <Header
          image="/src/assets/hackgt-logo.png"
          content="Sponsorship Portal"
          textAlign="center"
        />
      </Container>
    );
  }
}

export default connect()(Homepage);
