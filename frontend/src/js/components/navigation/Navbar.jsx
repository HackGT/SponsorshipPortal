import React from 'react';
import { Container, Header, Menu, Button } from 'semantic-ui-react';
import logo from '../../../assets/hackgt-logo.png';
import { TITLE } from '../../configs';

class Navbar extends React.Component {
  render() {
    // const page = this.props.location.pathname;
    const ui = this.props.ui;
    const auth = this.props.auth;
    const logOut = this.props.logOut;
    const isLoggedIn = auth.get('active');
    const syncStatus = ui.get('syncStatus');
    const logoutButton = (!isLoggedIn) ? false : (
      <Button
        onClick={() => { logOut(); }}
        basic
      >
        Sign out
      </Button>
    );
    // const push = this.props.push;
    return (
      <div style={{ top: 0, width: '100vw', backgroundColor: '#ffffff', borderBottom: '1px solid grey' }}>
        <Container>
          <Menu secondary>
            <Menu.Header>
              <Header
                image={logo}
                content={TITLE}
                textAlign="center"
              />
            </Menu.Header>
            {/* <Menu.Item
              position="left"
              name="pile"
            >
              <Button.Group>
                <Button primary={page === '/'} onClick={() => push('/')}>Home</Button>
                <Button primary={page === '/participant-profile'} onClick={() => push('/participant-profile')}>Workspace</Button>
                <Button primary={page === '/export'} onClick={() => push('/export')}>Export</Button>
              </Button.Group>
            </Menu.Item> */}
            <Menu.Item
              position="right"
            >
              {syncStatus}
              {logoutButton}
            </Menu.Item>
          </Menu>
        </Container>
      </div>
    );
  }
}

export default Navbar;
