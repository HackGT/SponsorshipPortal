import React from 'react';
import { Button, Header, Menu } from 'semantic-ui-react';

class Navbar extends React.Component {
  render() {
    const page = this.props.page;
    return (
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
          <Button.Group>
            <Button primary={page === 'HomePage'}>Home</Button>
            <Button primary={page === 'ParticipantProfilePage'}>Pile</Button>
          </Button.Group>
        </Menu.Item>
        <Menu.Item
          position="right"
          name="logout"
        />
      </Menu>
    );
  }
}

export default Navbar;
