import React from 'react';
import { connect } from 'react-redux';
import { Button, Container, Form, Grid, Header, Input, Menu, Table, Checkbox, Icon } from 'semantic-ui-react';

import Navbar from '../../components/navigation/Navbar';

class HomePage extends React.Component {
  render() {
    return (
      <Container>
        <Navbar page="HomePage" />
        <p>This page is under construction</p>
      </Container>
    );
  }
}

export default connect()(HomePage);
