import React from 'react';
import { Container, Button, Header, Menu } from 'semantic-ui-react';

class Navbar extends React.Component {
  render() {
    const page = this.props.location.pathname;
    const push = this.props.push;
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
            position="left"
            name="pile"
          >
            <Button.Group>
              <Button primary={page === '/'} onClick={() => push('/')}>Home</Button>
              <Button primary={page === '/participant-profile'} onClick={() => push('/participant-profile')}>Workspace</Button>
              <Button primary={page === '/export'} onClick={() => push('/export')}>Export</Button>
            </Button.Group>
          </Menu.Item>
          <Menu.Item
            position="right"
            name="logout"
          />
        </Menu>
      </Container>
    );
  }
}

export default Navbar;
