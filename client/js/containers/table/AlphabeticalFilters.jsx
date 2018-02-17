import React from 'react';
import { connect } from 'react-redux';
import { Button, Grid } from 'semantic-ui-react';
import * as actions from '../../actions/participants';

class AlphabeticalFilters extends React.Component {
  render() {
    const participants = this.props.participants;
    const showParticipant = this.props.showParticipant;
    const hideParticipant = this.props.hideParticipant;
    const letters = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'];

    return (
      <Grid>
        <Grid.Row>
          <Button.Group widths="13" fluid>
            {
              letters.slice(0, 13).map((letter) => {
                return (
                  <Button
                    basic
                    key={letter}
                    onClick={() => {
                      // show participants with name starting at this letter
                      participants.forEach((participant) => {
                        if (participant.get('name').toLowerCase().charAt(0) === letter) {
                          showParticipant(participant);
                        } else {
                          hideParticipant(participant);
                        }
                      });
                    }}
                  >
                    {letter.toUpperCase()}
                  </Button>
                );
              })
            }
          </Button.Group>
          <Button.Group width="13" fluid>
            {
              letters.slice(13, 26).map((letter) => {
                return (
                  <Button
                    basic
                    key={letter}
                    onClick={() => {
                      // show participants with name starting at this letter
                      participants.forEach((participant) => {
                        if (participant.get('name').toLowerCase().charAt(0) === letter) {
                          showParticipant(participant);
                        } else {
                          hideParticipant(participant);
                        }
                      });
                    }}
                  >
                    {letter.toUpperCase()}
                  </Button>
                );
              })
            }
          </Button.Group>
        </Grid.Row>
      </Grid>
    );
  }
}

export default connect((state) => {
  return {
    participants: state.get('participants'),
  };
}, (dispatch) => {
  return {
    showParticipant: (participant) => {
      dispatch(actions.showParticipant(participant));
    },
    hideParticipant: (participant) => {
      dispatch(actions.hideParticipant(participant));
    },
  };
})(AlphabeticalFilters);
