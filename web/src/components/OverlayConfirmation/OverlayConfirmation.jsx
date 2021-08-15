import React from 'react';
import PropTypes from 'prop-types';
import './OverlayConfirmation.scss';

const OverlayConfirmation = ({show, loading, title, okText, cancelText, onOk, onCancel}) => {
  const handleOverClicked = e => {
    e.preventDefault();
    if (e.target === e.currentTarget) {
      onCancel();
    }
  };

  return (
    <React.Fragment>
      {show && (
        <div className="overlay_confirmation">
          <div className="overlay_confirmation__content" onClick={handleOverClicked}>
            {loading ? (
              <div className="overlay_confirmation__loader" />
            ) : (
              <React.Fragment>
                <div className="overlay_confirmation__text">{title}</div>
                <div className="overlay_confirmation__btn overlay_confirmation__btn--ok" onClick={onOk}>{okText}</div>
                <div className="overlay_confirmation__btn overlay_confirmation__btn--cancel"
                     onClick={onCancel}>{cancelText}</div>
              </React.Fragment>
            )}
          </div>
        </div>
      )}
    </React.Fragment>
  );
};

OverlayConfirmation.propTypes = {
  show: PropTypes.bool,
  title: PropTypes.string,
  okText: PropTypes.string,
  cancelText: PropTypes.string,
  onOk: PropTypes.func,
  onCancel: PropTypes.func,
  loading: PropTypes.bool,
};

OverlayConfirmation.defaultProps = {
  show: false,
  loading: false,
  title: 'Are you sure?',
  okText: 'Ok',
  cancelText: 'Cancel',
  onOk: () => {
  },
  onCancel: () => {
  },
};

export default OverlayConfirmation;
