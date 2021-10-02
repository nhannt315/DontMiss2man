import React, { KeyboardEvent, useCallback } from 'react';
import {
  Paper as MuiPaper,
  Popper as MuiPopper,
  PopperProps as MuiPopperProps,
  PopoverProps as MuiPopoverProps,
} from '@material-ui/core';

export interface PopperProps {
  anchor?: HTMLElement | null;
  disablePaper?: boolean;
  disablePortal?: MuiPopperProps['disablePortal'];
  // modifiers={{ offset: { offset: '-76, 22' } }}
  modifiers?: MuiPopperProps['modifiers'];
  placement?: MuiPopperProps['placement'];
  popperOptions?: MuiPopperProps['popperOptions'];
  onClose?: MuiPopoverProps['onClose'];
  open: MuiPopperProps['open'];
}

// MaterialUIのPopoverだと、閉じるアクションと次のアクションが別のステップになってしまう
// 対してPopperだと、外側のクリックした時に、onCloseと次のアクションを同時に行えるが良い
// 代わりに、Escのclose処理を自前で書く必要がある
const Popper: React.FC<PopperProps> = ({
  children,
  anchor,
  disablePortal,
  modifiers,
  placement,
  popperOptions,
  onClose,
  open,
}) => {
  const handleKeyDown = useCallback(
    (e: KeyboardEvent) => {
      e.stopPropagation();
      if (e.key === 'Escape') {
        onClose?.(e, 'escapeKeyDown');
      }
    },
    [onClose]
  );

  return (
    <MuiPopper
      anchorEl={anchor}
      disablePortal={disablePortal}
      modifiers={modifiers}
      onKeyDown={handleKeyDown}
      open={open}
      placement={placement}
      popperOptions={popperOptions}
      // モーダルなど他の重なり要素との兼ね合いで調整
      style={{ zIndex: 9999 }}
    >
      {/* TrapFocusすると、ModalのなかでPopper使った時にエラー↓が出る */}
      {/* Uncaught RangeError: Maximum call stack size exceeded. */}
      {/* <TrapFocus
        open
        isEnabled={() => true}
        getDoc={() => anchor?.ownerDocument ?? document}
      > */}
      <MuiPaper
        style={{
          background: 'white',
          overflowX: 'hidden',
          overflowY: 'auto',
          outline: 'none',
        }}
        tabIndex={-1}
      >
        {children}
      </MuiPaper>
      {/* </TrapFocus> */}
    </MuiPopper>
  );
};

export default Popper;
