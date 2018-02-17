import { HOST } from '../configs';
import NotificationHelper from './NotificationHelper';
import store from '../index';
import { showPDFAtURLInModal } from '../actions/ui';

const PDFHelper = {};

/**
 * Retrieve Resume URL and execute callback(url)
 */
PDFHelper.findResumeURL = (resumeId, callback) => {
  fetch(`${HOST}/viewResume`, {
    method: 'POST',
    headers: new Headers({ 'Content-Type': 'application/json' }),
    body: JSON.stringify({
      resumeId,
      token: store.getState().get('auth').get('token'),
    }),
  }).then((response) => {
    if (response.ok) {
      return response.json();
    }
    throw new Error('POST /viewResume connection lost');
  }).then((json) => {
    if (!json.fileURL) {
      throw new Error('Invalid URL response');
    }

    callback(json.fileURL);
  }).catch(() => {
    NotificationHelper.showModalWithMessage('Connection lost. Please reload this page.');
  });
};

/**
 * Show Resume PDF files in a Modal
 */
PDFHelper.showResumeInModal = (resumeId) => {
  PDFHelper.findResumeURL(resumeId, (url) => {
    store.dispatch(showPDFAtURLInModal(url));
  });
};

/**
 * Show Resume in a new tab or window
 */
PDFHelper.showResumeInNewTab = (resumeId) => {
  PDFHelper.findResumeURL(resumeId, (url) => {
    /*
     * Try to load the pdf without being recogized as a pop up
     * This has different behaviors depending on browsers 
     */
    const isChrome = /Chrome/.test(navigator.userAgent) && /Google Inc/.test(navigator.vendor);
    if (isChrome) {
      // Should work on Chrome and Webkit
      // Use anchor element to avoid being recognized as pop up
      const link = document.createElement('a');
      link.setAttribute('href', url);
      link.setAttribute('target', '_blank');
      link.click();
    } else {
      // Works on Firefox
      window.open(url, '_blank');
    }
  });
};

export default PDFHelper;
