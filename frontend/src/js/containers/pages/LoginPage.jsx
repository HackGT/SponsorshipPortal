import React from 'react';
import { connect } from 'react-redux';
import { Button, Container, Segment, Form, Input } from 'semantic-ui-react';
import { logIn } from '../../actions/auth';

/**
 * A temporary auth solution. Will be replaced when the central auth service is completed
 */
class LoginPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      username: '',
      password: '',
    };
  }

  render() {
    return (
      <Container style={{ marginTop: '20vh' }}>
        <Segment>
          <h1>Welcome</h1>
          <h3>Please Log in with your Sponsorship Portal Credentials</h3>
        </Segment>
        <Segment inverted raised>
          <Form
            inverted
            onSubmit={() => {
              this.props.logIn(this.state.username, this.state.password);
              this.setState({
                username: '',
                password: '',
              });
            }}
          >
            <Form.Field>
              <Input
                placeholder="Username"
                value={this.state.username}
                onChange={(event, data) => {
                  this.setState({
                    username: data.value,
                  });
                }}
              />
            </Form.Field>
            <Form.Field>
              <Input
                placeholder="Password"
                type="password"
                value={this.state.password}
                onChange={(event, data) => {
                  this.setState({
                    password: data.value,
                  });
                }}
              />
            </Form.Field>
            <Button type="submit">Log In</Button>
          </Form>
        </Segment>
      </Container>
    );
  }
}

function mapDispatchToProps(dispatch) {
  return {
    logIn: (username, password) => dispatch(logIn(username, password)),
  };
}

export default connect(null, mapDispatchToProps)(LoginPage);
