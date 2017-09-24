import React from 'react';
import { connect } from 'react-redux';
import { Container } from 'semantic-ui-react';


class ExportPage extends React.Component {
  render() {
    return (
      <Container>
        <h2>Choose An Export Format</h2>
      </Container>
    );
  }
}

export default connect()(ExportPage);
