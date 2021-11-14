import React from 'react';
import {
  Checkbox as MuiCheckbox,
  CheckboxProps as MuiCheckboxProps,
  FormControlLabel,
  FormControlLabelProps,
} from '@material-ui/core';

interface CheckboxProps {
  className?: string;
  checked?: MuiCheckboxProps['checked'];
  disabled?: MuiCheckboxProps['disabled'];
  inputRef?: MuiCheckboxProps['inputRef'] | FormControlLabelProps['inputRef'];
  label?: FormControlLabelProps['label'];
  name?: MuiCheckboxProps['name'];
  onChange?: MuiCheckboxProps['onChange'];
}

export const Checkbox: React.FC<CheckboxProps> = ({
  className,
  checked = false,
  disabled = false,
  inputRef,
  label,
  name,
  onChange,
}) => {
  const renderCheckbox = (ref?: MuiCheckboxProps['inputRef']) => {
    return (
      <MuiCheckbox
        checked={checked}
        disabled={disabled}
        inputRef={ref}
        name={name}
        onChange={onChange}
        size="small"
        disableRipple
        disableFocusRipple
        disableTouchRipple
      />
    );
  };

  if (label === undefined) {
    return renderCheckbox(inputRef);
  }

  return (
    <FormControlLabel
      className={`${className}`}
      control={renderCheckbox()}
      label={label}
      inputRef={inputRef}
    />
  );
};
