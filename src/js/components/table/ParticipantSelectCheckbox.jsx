import React from 'react';
import { Button } from 'semantic-ui-react';

class ParticipantSelectCheckbox extends React.Component {
  render() {
    const checked = this.props.checked;
    const onClick = this.props.onClick;
    const label = checked ? 'Remove' : 'Select';
    const icon = checked ? 'user delete' : 'plus';
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
