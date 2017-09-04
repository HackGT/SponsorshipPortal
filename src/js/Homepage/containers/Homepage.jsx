import React from 'react';
import { connect } from 'react-redux';
import { Button, Container, Form, Grid, Header, Input, Menu, Table, Checkbox, Icon } from 'semantic-ui-react';

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
            position="center"
            name="pile"
          >
            <Button primary>Your Pile</Button>
          </Menu.Item>
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
                  label="HackGT Participant Search"
                  placeholder="Enter name or keywords"
                />
              </Form.Field>
            </Form>
          </Grid.Row>
          <Grid.Row columns="1">
            <Table compact celled definition>
              <Table.Header>
                <Table.Row>
                  <Table.HeaderCell />
                  <Table.HeaderCell>Name</Table.HeaderCell>
                  <Table.HeaderCell>School</Table.HeaderCell>
                  <Table.HeaderCell>E-mail address</Table.HeaderCell>
                  <Table.HeaderCell>Resume</Table.HeaderCell>
                  <Table.HeaderCell>Skills</Table.HeaderCell>
                </Table.Row>
              </Table.Header>

              <Table.Body>
                <Table.Row>
                  <Table.Cell collapsing>
                    <Checkbox slider />
                  </Table.Cell>
                  <Table.Cell>John Lilki</Table.Cell>
                  <Table.Cell>Georgia Tech</Table.Cell>
                  <Table.Cell>jhlilk22@yahoo.com</Table.Cell>
                  <Table.Cell><Button>View Resume</Button></Table.Cell>
                  <Table.Cell>Java, Python, C</Table.Cell>
                </Table.Row>
                <Table.Row>
                  <Table.Cell collapsing>
                    <Checkbox slider />
                  </Table.Cell>
                  <Table.Cell>Jamie Harington</Table.Cell>
                  <Table.Cell>Harvard</Table.Cell>
                  <Table.Cell>jamieharingonton@yahoo.com</Table.Cell>
                  <Table.Cell><Button>View Resume</Button></Table.Cell>
                  <Table.Cell>Go, Ruby, Haskell</Table.Cell>
                </Table.Row>
                <Table.Row>
                  <Table.Cell collapsing>
                    <Checkbox slider />
                  </Table.Cell>
                  <Table.Cell>Jill Lewis</Table.Cell>
                  <Table.Cell>Princeton</Table.Cell>
                  <Table.Cell>jilsewris22@yahoo.com</Table.Cell>
                  <Table.Cell><Button>View Resume</Button></Table.Cell>
                  <Table.Cell>SQL, Docker, Kubernetes</Table.Cell>
                </Table.Row>
              </Table.Body>

              <Table.Footer fullWidth>
                <Table.Row>
                  <Table.HeaderCell />
                  <Table.HeaderCell colSpan='4'>
                    <Button size='small'>Add to Pile</Button>
                    <Button size='small'>Add All to Pile</Button>
                  </Table.HeaderCell>
                </Table.Row>
              </Table.Footer>
            </Table>
          </Grid.Row>
        </Grid>
      </Container>
    );
  }
}

export default connect()(Homepage);
