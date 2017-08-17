import React from 'react';
import { connect } from 'react-redux';
import { Container, Form, Grid, Header, Input, Menu } from 'semantic-ui-react';

class Homepage extends React.Component {
  render() {
    return (
      <Container>
        <Menu secondary>
          <Menu.Header>
            <Header
              image="/src/assets/hackgt-logo.png"
              content="Sponsorship Portal"
              textAlign="center"
            />
          </Menu.Header>
          <Menu.Item
            position="right"
            name="logout"
          />
        </Menu>
        <Grid centered>
          <Grid.Row columns="1">
            <Form>
              <Form.Field>
                <Input
                  label="Search for HackGT participant"
                  placeholder="Search here"
                />
              </Form.Field>
            </Form>
          </Grid.Row>
        </Grid>
      </Container>
    );
  }
}

export default connect()(Homepage);
