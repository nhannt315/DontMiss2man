import React, {useState} from 'react';
import PropTypes from 'prop-types';
import {Button, Form, Icon, Input, Spin} from 'antd';
import i18n from '../../config/i18n';
import './RegisterForm.scss';

const RegisterForm = ({form, submit, className, loading}) => {

  const [confirmDirty, setConfirmDirty] = useState(false);

  const handleSubmit = e => {
    e.preventDefault();
    form.validateFields((err, values) => {
      if (!err) {
        submit(values);
      }
    });
  };

  const handleConfirmBlur = e => {
    const {value} = e.target;
    setConfirmDirty(confirmDirty || !!value);
  };

  const compareToFirstPassword = (rule, value, callback) => {
    if (value && value !== form.getFieldValue('password')) {
      callback(i18n.t('auth.password_not_match'));
    } else {
      callback();
    }
  };

  const validateToNextPassword = (rule, value, callback) => {
    if (value && confirmDirty) {
      form.validateFields(['confirm'], {force: true});
    }
    callback();
  };

  const {getFieldDecorator} = form;
  return (
    <Form onSubmit={handleSubmit} className={['auth_form', className].join(' ')}>
      <Spin spinning={loading}>
        <Form.Item label={i18n.t('auth.mail_address')}>
          {getFieldDecorator('email', {
            rules: [{required: true, message: i18n.t('auth.email_required')}],
          })(
            <Input type="email" placeholder={i18n.t('auth.email')} />,
          )}
        </Form.Item>
        <Form.Item label={i18n.t('auth.password')}>
          {getFieldDecorator('password', {
            rules: [{required: true, message: i18n.t('auth.password_required')}, {validator: validateToNextPassword}],
          })(
            <Input.Password
              prefix={<Icon type="lock" style={{color: 'rgba(0,0,0,.25)'}} />}
              placeholder={i18n.t('auth.password')}
            />,
          )}
        </Form.Item>
        <Form.Item label={i18n.t('auth.password_confirm')}>
          {getFieldDecorator('password_confirm', {
            rules: [{required: true, message: i18n.t('auth.password_required')}, {validator: compareToFirstPassword}],
          })(
            <Input.Password
              prefix={<Icon type="lock" style={{color: 'rgba(0,0,0,.25)'}} />}
              placeholder={i18n.t('auth.password_confirm')}
              onBlur={handleConfirmBlur}
            />,
          )}
        </Form.Item>
        <Button type="primary" htmlType="submit" className="auth_form-button">
          {i18n.t('auth.register')}
        </Button>
      </Spin>
    </Form>
  );
};

RegisterForm.propTypes = {
  className: PropTypes.string,
  form: PropTypes.object,
  submit: PropTypes.func,
  loading: PropTypes.bool,
};

RegisterForm.defaultProps = {
  className: '',
  form: {},
  submit: () => {
  },
  loading: false,
};

export default Form.create()(RegisterForm);
