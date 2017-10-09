/**
 * The current snapshot mechanism is temporary and inefficient. This should be optimized when the data gets big.
 */
const SyncHelper = {};

SyncHelper.getSnapshotFromState = (participantsList) => {
  return JSON.stringify(participantsList);
};

SyncHelper.getStateFromSnapshot = (snapshot) => {
  return JSON.parse(snapshot);
};

export default SyncHelper;
