import React, { useMemo, useState } from 'react';

interface CheckboxValue {
  label: string;
  value: string;
  checked: boolean;
}

interface Props {
  checkboxes: CheckboxValue[];
  onChange: () => void;
}

export const CheckboxGroup: React.FC<Props> = ({ checkboxes }) => {
  const [parentCheckboxChecked, setParentCheckboxChecked] = useState<boolean>(
    false
  );

  const updateParentWithChildren = useMemo(() => {
    let allChecked = false;
    for (let i = 0; i < checkboxes.length; i += 1) {
      if (checkboxes[i].checked) {
        allChecked = true;
      } else {
        allChecked = false;
        break;
      }
    }
    setParentCheckboxChecked(allChecked);
  }, [checkboxes]);

  return <div>Check box Groupd</div>;
};
