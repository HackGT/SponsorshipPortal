import React from 'react';
import ConnectedNavbar from './containers/navigation/ConnectedNavbar';
import GlobalLoader from './containers/global/GlobalLoader';

export default function App({ children, location }) {
  return (
    <div>
      <ConnectedNavbar location={location} />
      <GlobalLoader />
      <div style={{ marginTop: '1.5em' }}>{children}</div>
    </div>
  );
}
