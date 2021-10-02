import React from 'react';
import PropTypes from 'prop-types';

import './BuildingList.scss';
import BuildingItem from './BuildingItem';

const BuildingList = ({buildingList, history, userData, handleFavoriteAction, isAuthenticated}) => {
  return (
    <div>
      {buildingList.map(item => {
        return <BuildingItem
          key={item.id} history={history} item={item} userData={userData}
          isAuthenticated={isAuthenticated} handleFavoriteAction={handleFavoriteAction}
        />;
      })}
    </div>
  );
};

BuildingList.propTypes = {
  buildingList: PropTypes.array,
  history: PropTypes.object.isRequired,
  userData: PropTypes.object,
  isAuthenticated: PropTypes.bool,
  handleFavoriteAction: PropTypes.func.isRequired,
};

BuildingList.defaultProps = {
  buildingList: [],
  userData: {},
  isAuthenticated: false,
};

export default BuildingList;
