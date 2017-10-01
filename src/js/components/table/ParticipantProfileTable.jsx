import React from 'react';
import { Button, Container, Grid, Table } from 'semantic-ui-react';

import ParticipantSelectCheckbox from './ParticipantSelectCheckbox';
import ParticipantProfileSearchBar from '../../containers/table/ParticipantProfileSearchBar';
import ParticipantProfileFilters from '../../containers/table/ParticipantProfileFilters';

class ParticipantProfileTable extends React.Component {
  render() {
    const participants = this.props.participants;
    const selectParticipant = this.props.selectParticipant;
    const unSelectParticipant = this.props.unSelectParticipant;
    // Thinking about letting user to "swipe right/swipe left" and blacklist participants
    // const showParticipant = this.props.showParticipant;
    // const hideParticipant = this.props.hideParticipant;
    return (
      <Container>
        <Grid centered>
          <Grid.Row columns="2">
            <Grid.Column>
              <ParticipantProfileSearchBar />
            </Grid.Column>
            <Grid.Column>
              <ParticipantProfileFilters />
            </Grid.Column>
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
                  <Table.HeaderCell>Tags</Table.HeaderCell>
                </Table.Row>
              </Table.Header>

              <Table.Body>
                {
                  participants.map((participant) => {
                    if (participant.get('isDisplaying')) {
                      return (
                        <Table.Row key={participant.get('id')}>
                          <Table.Cell collapsing>
                            <ParticipantSelectCheckbox
                              checked={participant.get('isSelected')}
                              onClick={() => {
                                if (!participant.get('isSelected')) {
                                  // if checkbox state is false
                                  selectParticipant(participant);
                                } else {
                                  unSelectParticipant(participant);
                                }
                              }}
                            />
                            {/* <Checkbox
                              slider
                              checked={participant.get('isSelected')}
                              onChange={() => {
                                if (!participant.get('isSelected')) {
                                  // if checkbox state is false
                                  selectParticipant(participant);
                                } else {
                                  unSelectParticipant(participant);
                                }
                              }}
                            /> */}
                          </Table.Cell>
                          <Table.Cell>{participant.get('name') || 'N/A'}</Table.Cell>
                          <Table.Cell>{participant.get('school') || 'N/A'}</Table.Cell>
                          <Table.Cell>{participant.get('email') || 'N/A'}</Table.Cell>
                          <Table.Cell><Button>View Resume</Button></Table.Cell>
                          <Table.Cell>Java, Python</Table.Cell>
                          <Table.Cell>None</Table.Cell>
                        </Table.Row>
                      );
                    }
                    return null;
                  })
                }
              </Table.Body>

              <Table.Footer fullWidth>
                <Table.Row>
                  <Table.HeaderCell />
                  <Table.HeaderCell colSpan="4">
                    <Button
                      size="small"
                      onClick={
                        () => {
                          participants
                            .filter((participant) => { return participant.get('isDisplaying'); })
                            .forEach((participant) => {
                              selectParticipant(participant);
                            });
                        }
                      }
                    >
                      Select All Displaying
                    </Button>
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

export default ParticipantProfileTable;
