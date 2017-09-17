import React from 'react';
import ConnectedNavbar from './containers/navigation/ConnectedNavbar';

export default function App({ children, location }) {
  return (
    <div>
      <ConnectedNavbar location={location} />
      <div style={{ marginTop: '1.5em' }}>{children}</div>
    </div>
  );
}
