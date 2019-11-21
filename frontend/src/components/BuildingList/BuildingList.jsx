import React from 'react';
import PropTypes from 'prop-types';

import './BuildingList.scss';
import BuildingItem from './BuildingItem';

const BuildingList = ({buildingList}) => {
  return (
    <div>
      {buildingList.map(item => {
        return <BuildingItem key={item.id} item={item} />;
      })}
    </div>
  );
};

BuildingList.propTypes = {
  buildingList: PropTypes.array,
};

BuildingList.defaultProps = {
  buildingList: [],
};

export default BuildingList;
