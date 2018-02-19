import React from 'react';
import { connect } from 'react-redux';
import { Modal, Button, Icon } from 'semantic-ui-react';
import * as ACTION_TYPE from '../../actionTypes';

function GlobalModal({ active, message, close }) {
  return (
    <Modal
      basic
      size="small"
      open={active}
      closeOnDimmerClick={false}
    >
      <Modal.Content>
        <p>{ message }</p>
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
    active: state.get('ui').get('modalActive'),
    message: state.get('ui').get('modalMessage'),
  };
}, (dispatch) => {
  return {
    close: () => dispatch({
      type: ACTION_TYPE.MODAL_OFF,
    }),
  };
})(GlobalModal);
