import React from 'react';
import PropTypes from 'prop-types';
import * as _ from 'lodash';
import {Skeleton} from 'antd';

const ListPlaceholder = ({size, itemCount}) => {
  const list = [];
  _.times(itemCount, i => {
    if (size === 'small') {
      list.push(<Skeleton key={i} active loading avatar paragraph={{rows: 2}} />);
    } else if (size === 'medium') {
      list.push(<Skeleton key={i} active loading avatar paragraph={{rows: 4}} />);
    }
  });
  return (
    <React.Fragment>
      {list}
    </React.Fragment>
  );
};

ListPlaceholder.propTypes = {
  size: PropTypes.oneOf(['small', 'medium']),
  itemCount: PropTypes.number,
};

ListPlaceholder.defaultProps = {
  size: 'medium',
  itemCount: 10,
};

export default ListPlaceholder;
