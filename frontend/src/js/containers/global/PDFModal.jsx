import React from 'react';
import { connect } from 'react-redux';
import { Document, Page } from 'react-pdf/build/entry.webpack';
import { Modal, Button, Icon } from 'semantic-ui-react';
import * as ACTION_TYPE from '../../actionTypes';


function PDFModal({ active, url, close }) {
  return (
    <Modal
      size="fullscreen"
      open={active}
      closeOnDimmerClick={false}
    >
      <Modal.Content>
        <Document file={url}>
          <Page pageNumber={1} />
        </Document>
      </Modal.Content>
      <Modal.Actions>
        <Button
          basic
          color="red"
          inverted
          onClick={() => {
            close();
          }}
        >
          <Icon name="remove" /> Close
        </Button>
      </Modal.Actions>
    </Modal>
  );
}

export default connect((state) => {
  return {
    active: state.get('ui').get('pdfModalActive'),
    url: state.get('ui').get('pdfUrl'),
  };
}, (dispatch) => {
  return {
    close: () => dispatch({
      type: ACTION_TYPE.UI_PDF_MODAL_OFF,
    }),
  };
})(PDFModal);
