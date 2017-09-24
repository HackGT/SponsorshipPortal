import React from 'react';
import { Button } from 'semantic-ui-react';

class ParticipantSelectCheckbox extends React.Component {
  render() {
    const checked = this.props.checked;
    const onClick = this.props.onClick;
    const label = checked ? 'Select' : 'Remove';
    const icon = checked ? 'plus' : 'user delete';
    return (
      <Button
        content={label}
        icon={icon}
        onClick={() => {
          onClick();
        }}
      />
    );
  }
}

export default ParticipantSelectCheckbox;
