import React from 'react';
import { Form, Input } from 'semantic-ui-react';

/**
 * Search bar using local state. Fetches for full-text search result upon submission.
 */
class ParticipantProfileSearchBar extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      content: '',
    };
  }
  render() {
    return (
      <Form
        onSubmit={() => {
          // Works when clicking "Search" action or pressing enter
          // TODO: Send request to backend and fetch search result
          console.log(this.state.content); // eslint-disable-line no-console
          this.setState({
            content: '',
          });
        }}
      >
        <Form.Field>
          <Input
            action="Search"
            placeholder="Search name or keywords"
            value={this.state.content}
            onChange={(event, data) => {
              this.setState({
                content: data.value,
              });
            }}
          />
        </Form.Field>
      </Form>
    );
  }
}

export default ParticipantProfileSearchBar;
