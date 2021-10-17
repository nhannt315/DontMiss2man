import React from 'react';
import Select from './index';

export default {
  title: 'components/Select',
  component: Select,
};

export const Basic: React.FC = () => {
  return (
    <>
      <div>
        <div>Select component</div>
        <Select
          className="w-20 h-48"
          options={[
            { label: '5万円', value: '5' },
            { label: '6万円', value: '6' },
            { label: '7万円', value: '7' },
            { label: '8万円', value: '8' },
            { label: '8万円', value: '8' },
            { label: '8万円', value: '8' },
            { label: '8万円', value: '8' },
            { label: '8万円', value: '8' },
            { label: '8万円', value: '8' },
            { label: '8万円', value: '8' },
            { label: '8万円', value: '8' },
          ]}
          onChange={(value) => alert(`new value: ${value}`)}
        />
      </div>
    </>
  );
};
