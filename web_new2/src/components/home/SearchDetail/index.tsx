import React from 'react';

interface IProps {
  className?: string;
}

const SearchDetail: React.FC<IProps> = ({ className }) => {
  return <div className={className}>Search Detail</div>;
};

export default SearchDetail;
