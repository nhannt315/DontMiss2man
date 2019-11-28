import React from 'react';
import PropTypes from 'prop-types';
import * as _ from 'lodash';
import {Skeleton} from 'antd';

const ListPlaceholder = ({size, itemCount, avatar}) => {
  const list = [];
  _.times(itemCount, i => {
    if (size === 'small') {
      list.push(<Skeleton key={i} active loading avatar={avatar} paragraph={{rows: 2}} />);
    } else if (size === 'medium') {
      list.push(<Skeleton key={i} active loading avatar={avatar} paragraph={{rows: 4}} />);
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
  avatar: PropTypes.bool,
};

ListPlaceholder.defaultProps = {
  size: 'medium',
  itemCount: 10,
  avatar: false,
};

export default ListPlaceholder;
