import React from 'react';
import { connect } from 'react-redux';
import { Container, Button } from 'semantic-ui-react';
import ExportHelper from '../../service/ExportHelper';


class ExportPage extends React.Component {
  render() {
    return (
      <Container style={{ marginBottom: '40px' }}>
        <h2>Choose An Export Format</h2>
        <p>
          <Button
            secondary
            onClick={() => {
              ExportHelper.exportCSV();
            }}
          >
            Download
          </Button>
           Export All Selected Candidates into CSV
        </p>
      </Container>
    );
  }
}

export default connect()(ExportPage);
