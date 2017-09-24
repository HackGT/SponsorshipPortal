import React from 'react';
import { connect } from 'react-redux';
import { Container, Button } from 'semantic-ui-react';


class ExportPage extends React.Component {
  render() {
    return (
      <Container>
        <h2>Choose An Export Format</h2>
        <p>
          <Button secondary>Download</Button>
           Export All Selected Candidates into CSV
        </p>
      </Container>
    );
  }
}

export default connect()(ExportPage);
