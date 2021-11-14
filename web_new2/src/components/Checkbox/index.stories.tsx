import React, { useState } from 'react';
import { Checkbox } from './index';

export default {
  title: 'components/Checkbox',
  component: Checkbox,
};

export const Basic: React.FC = () => {
  const [checked, setChecked] = useState(false);

  return (
    <>
      <div>
        <div>Single checkbox section</div>
        <Checkbox
          checked={checked}
          onChange={() => setChecked((prev) => !prev)}
        />
        <Checkbox
          label="チェックボックス"
          checked={checked}
          onChange={() => setChecked((prev) => !prev)}
        />
        <Checkbox
          label={<span style={{ fontWeight: 700 }}>オリジナルラベル</span>}
          checked={checked}
          onChange={() => setChecked((prev) => !prev)}
        />
      </div>
      <div className="mt-10">
        <div>Checkbox group section</div>
      </div>
    </>
  );
};
