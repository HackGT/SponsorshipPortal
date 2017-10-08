import React from 'react';
import ConnectedNavbar from './containers/navigation/ConnectedNavbar';
import GlobalLoader from './containers/global/GlobalLoader';

export default function App({ children, location }) {
  return (
    <div>
      <ConnectedNavbar location={location} />
      <GlobalLoader />
      <div style={{ paddingTop: '100px' }}>{children}</div>
    </div>
  );
}
