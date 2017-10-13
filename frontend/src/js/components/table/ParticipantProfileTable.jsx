import React from 'react';
import { Button, Container, Grid, Table, Dimmer, Loader, Segment } from 'semantic-ui-react';

import ParticipantSelectCheckbox from './ParticipantSelectCheckbox';
import ParticipantProfileSearchBar from '../../containers/table/ParticipantProfileSearchBar';
import ParticipantProfileFilters from '../../containers/table/ParticipantProfileFilters';
import AlphabeticalFilters from '../../containers/table/AlphabeticalFilters';

import PDFHelper from '../../service/PDFHelper';
import SyncHelper from '../../service/SyncHelper';
import ExportHelper from '../../service/ExportHelper';

class ParticipantProfileTable extends React.Component {
  render() {
    const participants = this.props.participants;
    const selectParticipant = this.props.selectParticipant;
    const unSelectParticipant = this.props.unSelectParticipant;
    const loaderActive = this.props.loaderActive;
    // Thinking about letting user to "swipe right/swipe left" and blacklist participants
    // const showParticipant = this.props.showParticipant;
    // const hideParticipant = this.props.hideParticipant;
    return (
      <Container>
        <Dimmer.Dimmable as={Segment} blurring dimmed={loaderActive}>
          <Dimmer
            active={loaderActive}
          >
            <Loader size="huge">Loading</Loader>
          </Dimmer>
          <Grid centered>
            <Grid.Row columns={1}>
              <Grid.Column>
                <ParticipantProfileSearchBar />
              </Grid.Column>
            </Grid.Row>
            <Grid.Row>
              <Grid.Column width={10}>
                <ParticipantProfileFilters />
              </Grid.Column>
              <Grid.Column width={6}>
                <Button
                  secondary
                  onClick={() => {
                    ExportHelper.exportCSV();
                  }}
                >
                  Export Selections to CSV
                </Button>
              </Grid.Column>
            </Grid.Row>
            <Grid.Row>
              <Grid.Column>
                <AlphabeticalFilters />
              </Grid.Column>
            </Grid.Row>
            <Grid.Row columns={1}>
              <Table compact celled definition>
                <Table.Header>
                  <Table.Row>
                    <Table.HeaderCell />
                    <Table.HeaderCell>Name</Table.HeaderCell>
                    <Table.HeaderCell>E-mail address</Table.HeaderCell>
                    <Table.HeaderCell>Resume</Table.HeaderCell>
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
                                  // request auto save
                                  SyncHelper.requestSync();
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
                            <Table.Cell>{participant.get('email') || 'N/A'}</Table.Cell>
                            <Table.Cell>
                              <Button
                                onClick={() => {
                                  PDFHelper.showResumeInNewTab(participant.get('resumeId'));
                                }}
                              >
                                View Resume
                              </Button>
                            </Table.Cell>
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
                      {/* <Button
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
                      </Button> */}
                    </Table.HeaderCell>
                  </Table.Row>
                </Table.Footer>
              </Table>
            </Grid.Row>
          </Grid>
        </Dimmer.Dimmable>
      </Container>
    );
  }
}

export default ParticipantProfileTable;
