import React from 'react';
import PropTypes from 'prop-types';

import './BuildingList.scss';
import BuildingItem from './BuildingItem';

const BuildingList = ({buildingList, history}) => {
  return (
    <div>
      {buildingList.map(item => {
        return <BuildingItem key={item.id} history={history} item={item} />;
      })}
    </div>
  );
};

BuildingList.propTypes = {
  buildingList: PropTypes.array,
  history: PropTypes.object.isRequired,
};

BuildingList.defaultProps = {
  buildingList: [],
};

export default BuildingList;
