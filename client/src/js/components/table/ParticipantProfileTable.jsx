import React from 'react';
import { Button, Container, Grid, Table, Dimmer, Loader, Segment } from 'semantic-ui-react';
import { Document, Page } from 'react-pdf';
import ReactTable from 'react-table';

import ParticipantSelectCheckbox from './ParticipantSelectCheckbox';
import ParticipantProfileSearchBar from '../../containers/table/ParticipantProfileSearchBar';
import ParticipantProfileFilters from '../../containers/table/ParticipantProfileFilters';

import PDFHelper from '../../service/PDFHelper';
import SyncHelper from '../../service/SyncHelper';
import ExportHelper from '../../service/ExportHelper';

class ParticipantProfileTable extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      currentlyViewingResumeType: null,
      currentlyViewingResumeURL: null,
      currentlyViewingParticipant: null,
    };
  }

  render() {
    const participants = this.props.participants;
    const selectParticipant = this.props.selectParticipant;
    const unSelectParticipant = this.props.unSelectParticipant;
    const loaderActive = this.props.loaderActive;
    // Thinking about letting user to "swipe right/swipe left" and blacklist participants
    // const showParticipant = this.props.showParticipant;
    // const hideParticipant = this.props.hideParticipant;

    // eslint-disable-next-line
    const leftColumnOld = (
      <Container fluid>
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
                              <Button.Group>
                                <Button
                                  positive
                                  onClick={() => {
                                    PDFHelper.showResumeInModal(participant.get('resumeId'));
                                  }}
                                >
                                  Read
                                </Button>
                                <Button.Or />
                                <Button
                                  onClick={() => {
                                    PDFHelper.showResumeInNewTab(participant.get('resumeId'));
                                  }}
                                >
                                  Download
                                </Button>
                              </Button.Group>
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

    const leftColumn = (
      <Container fluid>
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
              <ReactTable
                style={{ width: '100%' }}
                filterable
                defaultPageSize={12}
                showPaginationTop
                data={participants.toArray().filter((participant) => {
                  return participant.get('isDisplaying');
                })}
                columns={[
                  {
                    id: 'name',
                    Header: 'Name',
                    accessor: d => d.get('name'),
                  },
                  {
                    id: 'email',
                    Header: 'Email',
                    accessor: d => d.get('email'),
                  },
                  {
                    id: 'resume',
                    Header: 'Resume',
                    accessor: d => d,
                    Cell: row => (
                      <Button.Group>
                        <Button
                          positive
                          onClick={() => {
                            // PDFHelper.showResumeInModal(row.value.get('resumeId'));
                            PDFHelper.findResumeURL(row.value.get('resumeId'), (url) => {
                              this.setState({
                                currentlyViewingResumeType: 'pdf',
                                currentlyViewingParticipant: row.value,
                                currentlyViewingResumeURL: url,
                              });
                            });
                          }}
                        >
                          Read
                        </Button>
                        <Button.Or />
                        <Button
                          onClick={() => {
                            PDFHelper.showResumeInNewTab(row.value.get('resumeId'));
                          }}
                        >
                          Download
                        </Button>
                      </Button.Group>
                    ),
                  },
                  {
                    id: 'select',
                    Header: 'Action',
                    accessor: d => d,
                    Cell: (row) => {
                      const participant = row.value;
                      return (
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
                      );
                    },
                  },
                ]}
                // getTdProps={(state, rowInfo) => {
                //   if (rowInfo && rowInfo.original) {
                //     return {
                //       onClick: (e, handleOriginal) => {
                //         console.log('Row Clicked:', rowInfo.original);
                //         // React-Table uses onClick internally to trigger
                //         // events like expanding SubComponents and pivots.
                //         // By default a custom 'onClick' handler will override this functionality.
                //         // If you want to fire the original onClick handler, call the
                //         // 'handleOriginal' function.
                //         if (handleOriginal) {
                //           handleOriginal();
                //         }
                //       },
                //       style: {
                //         height: '3.7vh',
                //         background: rowInfo.original.onFocus ? 'teal' : 'white',
                //       },
                //     };
                //   }
                //   return {};
                // }}
              />
            </Grid.Row>
          </Grid>
        </Dimmer.Dimmable>
      </Container>
    );

    const rightColumn = (
      <div style={{ overflowY: 'scroll', height: '95vh' }}>
        {
          this.state.currentlyViewingResumeURL && this.state.currentlyViewingResumeType === 'pdf' ? (
            <Document
              file={this.state.currentlyViewingResumeURL}
            >
              <Page pageNumber={1} scale={1.5} />
            </Document>
          ) : (
            <div />
          )
        }
        {
          this.state.currentlyViewingResumeURL && this.state.currentlyViewingResumeType === 'doc' ? (
            <iframe title="word-viewer" src={`"http://docs.google.com/gview?url=${this.state.currentlyViewingResumeURL}&embedded=true"`} />
          ) : (
            <div />
          )
        }
      </div>
    );

    return (
      <Grid>
        <Grid.Row columns={2} only="computer tablet">
          <Grid.Column>{leftColumn}</Grid.Column>
          <Grid.Column>{rightColumn}</Grid.Column>
        </Grid.Row>
        <Grid.Row columns={1} only="mobile">
          {leftColumn}
        </Grid.Row>
        <Grid.Row columns={1} only="mobile">
          {rightColumn}
        </Grid.Row>
      </Grid>
    );
  }
}

export default ParticipantProfileTable;
