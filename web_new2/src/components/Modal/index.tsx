import React from 'react';
import {
  Dialog,
  DialogContent,
  DialogProps,
  DialogTitle,
  DialogActions,
  Modal as MuiModal,
  ModalProps as MuiModalProps,
} from '@material-ui/core';

export const ModalHeader: React.FC = ({ children }) => {
  return <DialogTitle>{children}</DialogTitle>;
};

export const ModalContent = DialogContent;

export const ModalActions = DialogActions;

interface ModalProps {
  fullScreen?: DialogProps['fullScreen'];
  fullwidth?: DialogProps['fullWidth'];
  onEntering?: DialogProps['onEntering'];
  onClose?: DialogProps['onClose'];
  open: boolean;
  size?: DialogProps['maxWidth'];
  transparent?: boolean;
  disablePortal?: boolean;
}

export const Modal: React.FC<ModalProps> = ({
  children,
  fullScreen = false,
  fullwidth = false,
  open,
  onEntering,
  onClose,
  size = 'md',
  transparent = false,
  disablePortal = false,
}) => {
  return (
    <Dialog
      fullScreen={fullScreen}
      open={open}
      onEntering={onEntering}
      onClose={onClose}
      maxWidth={size}
      fullWidth={fullwidth}
      PaperProps={{
        elevation: transparent ? 0 : 3,
      }}
      disablePortal={disablePortal}
    >
      {children}
    </Dialog>
  );
};

interface ModalBaseProps {
  children: MuiModalProps['children'];
  onClose?: MuiModalProps['onClose'];
  open: MuiModalProps['open'];
}

export const ModalBase: React.FC<ModalBaseProps> = ({
  children,
  onClose,
  open,
}) => {
  return (
    <MuiModal onClose={onClose} open={open}>
      {children}
    </MuiModal>
  );
};
