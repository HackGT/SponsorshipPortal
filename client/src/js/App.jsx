import React from 'react';
import { connect } from 'react-redux';
import ConnectedNavbar from './containers/navigation/ConnectedNavbar';
import GlobalLoader from './containers/global/GlobalLoader';
import GlobalModal from './containers/global/GlobalModal';
import PDFModal from './containers/global/PDFModal';
import LoginPage from './containers/pages/LoginPage';

function App({ children, location, loggedIn }) {
  return (
    <div>
      <ConnectedNavbar location={location} />
      <GlobalLoader />
      <GlobalModal />
      <PDFModal />
      {
        loggedIn ? <div>{ children }</div> : <LoginPage />
      }
    </div>
  );
}

export default connect((state) => {
  return {
    loggedIn: state.get('auth').get('active'),
  };
})(App);
