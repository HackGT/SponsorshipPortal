import React from 'react';
import { Form, Input } from 'semantic-ui-react';

class ParticipantProfileSearchBar extends React.Component {
  render() {
    return (
      <Form>
        <Form.Field>
          <Input
            label="HackGT Participant Search"
            placeholder="Enter name or keywords"
          />
        </Form.Field>
      </Form>
    );
  }
}

export default ParticipantProfileSearchBar;
