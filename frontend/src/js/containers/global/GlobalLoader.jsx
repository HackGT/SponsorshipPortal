import React from 'react';
import { connect } from 'react-redux';
import { Dimmer, Loader } from 'semantic-ui-react';

function GlobalLoader({ active }) {
  return (
    <Dimmer
      active={active}
    >
      <Loader>Loading</Loader>
    </Dimmer>
  );
}

export default connect((state) => {
  return {
    active: state.get('ui').get('loader'),
  };
})(GlobalLoader);
