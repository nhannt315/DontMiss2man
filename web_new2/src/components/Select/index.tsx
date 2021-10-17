import React, { useMemo, useRef, useState } from 'react';
import { ChevronUpIcon, ChevronDownIcon } from '@heroicons/react/solid';
import Popper from 'src/components/Popper';
import ClickAwayListener from '../ClickAwayListener';

export interface SelectOption {
  label: string | number;
  value: string | number;
}

interface Props {
  className?: string;
  options: SelectOption[];
  value?: string | number;
  onChange?: (newValue: string | number) => void;
}

const findOptionByValue = (
  options: SelectOption[],
  value?: string | number
): SelectOption => {
  const index = options.findIndex((item) => item.value === value);
  return options[index];
};

const Select: React.FC<Props> = ({ className, options, value, onChange }) => {
  const [currentValue, setCurrentValue] = useState<string | number>(
    value ?? options[0].value
  );
  const mainElement = useRef<HTMLDivElement>(null);
  const [anchor, setAnchor] = useState<HTMLElement | null>(null);

  const popperWidth: number = useMemo(() => {
    if (mainElement.current == null) return 0;
    return mainElement.current.offsetWidth;
  }, [mainElement]);

  return (
    <ClickAwayListener onClickAway={() => setAnchor(null)}>
      <div className={`${className}`} ref={mainElement}>
        <button
          className="border bg-white border-gray-200 hover:boder-blue-200 rounded px-1 flex flex-row items-center w-full"
          onClick={(e) => {
            setAnchor(anchor === null ? e.currentTarget : null);
          }}
        >
          <span className="flex-1 text-base py-1">
            {findOptionByValue(options, currentValue)?.label}
          </span>
          <span>
            {anchor === null ? (
              <ChevronDownIcon className="w-4" />
            ) : (
              <ChevronUpIcon className="w-4" />
            )}
          </span>
        </button>
        <Popper anchor={anchor} open={anchor !== null} placement="bottom-start">
          <div
            className={`flex flex-col h-28 overflow-y-scroll ${className} `}
            style={{ width: `${popperWidth}px` }}
          >
            {options.map((option) => (
              <button
                className="flex flex-row px-2 py-1 hover:bg-gray-200 text-center"
                key={option.value}
                onClick={() => {
                  setCurrentValue(option.value);
                  if (onChange) onChange(option.value);
                  setAnchor(null);
                }}
              >
                <div className="flex-1 text-base">{option.label}</div>
                <div className="w-4" />
              </button>
            ))}
          </div>
        </Popper>
      </div>
    </ClickAwayListener>
  );
};

export default Select;
