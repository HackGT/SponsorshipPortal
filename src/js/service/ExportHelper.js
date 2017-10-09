import store from '../index';

const json2csv = require('json2csv');

const ExportHelper = {};

ExportHelper.participantsListToCSVHelper = (participantsList) => {
  const fields = ['name', 'school', 'email'];

  // Filter and Normalize participantsList to fit the API
  const selectedParticipants = [];

  participantsList.forEach((participant) => {
    if (participant.get('isSelected')) {
      selectedParticipants.push({
        name: participant.get('name'),
        school: participant.get('school'),
        email: participant.get('email'),
      });
    }
  });

  // Return CSV
  try {
    const result = json2csv({ data: selectedParticipants, fields });
    return result;
  } catch (err) {
    // Errors are thrown for bad options, or if the data is empty and no fields are provided.
    // Be sure to provide fields if it is possible that your data array will be empty.
    console.error(err);
  }

  return null;
};

ExportHelper.downloadCSVHelper = (contentString) => {
  const encodedUri = encodeURI('data:text/csv;charset=utf-8,' + contentString);
  window.open(encodedUri);
};

ExportHelper.exportCSV = () => {
  const participantsList = store.getState().get('participants');
  ExportHelper.downloadCSVHelper(ExportHelper.participantsListToCSVHelper(participantsList));
};

export default ExportHelper;
