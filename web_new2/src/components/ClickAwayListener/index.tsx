import React from 'react';
import {
  ClickAwayListener as MuiClickAwayListener,
  ClickAwayListenerProps as MuiClickAwayListenerProps,
} from '@material-ui/core';

interface Props {
  onClickAway: MuiClickAwayListenerProps['onClickAway'];
}

const ClickAwayListener: React.FC<Props> = ({ children, onClickAway }) => {
  return (
    <MuiClickAwayListener onClickAway={onClickAway}>
      {children}
    </MuiClickAwayListener>
  );
};

export default ClickAwayListener;
