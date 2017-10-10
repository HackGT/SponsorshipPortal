import { HOST } from '../configs';
import NotificationHelper from './NotificationHelper';

const PDFHelper = {};

PDFHelper.showResumeInNewTab = (resumeId) => {
  fetch(`${HOST}/viewResume`, {
    token: store.getState().get('auth').get('token'),
    method: 'POST',
    resumeId,
  }).then((response) => {
    if (response.ok) {
      return response.json();
    }
    throw new Error('POST /viewResume connection lost');
  }).then((json) => {
    if (!json.fileURL) {
      throw new Error('Invalid URL response');
    }
    window.open(json.fileURL, '_blank'); // open pdf in a new tab / window / popup depending on browser settings
  }).catch(() => {
    NotificationHelper.showModalWithMessage('Connection lost. Please reload this page.');
  });
};

export default PDFHelper;
