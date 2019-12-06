import React from 'react';
import {Form, Icon, Input, Button, Checkbox} from 'antd';
import PropTypes from 'prop-types';
import {Link} from 'react-router-dom';
import i18n from '../../config/i18n';
import './LoginForm.scss';
import Logo from '../../assets/images/logo.png';

const LoginForm = ({form, className, styles, submit}) => {

  const handleSubmit = e => {
    e.preventDefault();
    form.validateFields((err, values) => {
      if (!err) {
        submit(values);
      }
    });
  };

  const {getFieldDecorator} = form;
  return (
    <Form onSubmit={handleSubmit} className={['auth_form', className].join(' ')} style={styles}>
      <div className="logo">
        <Link to="/">
          <img src={Logo} alt="Logo" />
          <span>DM2M</span>
        </Link>
      </div>
      <Form.Item>
        {getFieldDecorator('email', {
          rules: [{required: true, message: i18n.t('auth.email_required')}],
        })(
          <Input
            prefix={<Icon type="user" style={{color: 'rgba(0,0,0,.25)'}} />}
            type="email"
            placeholder={i18n.t('auth.email')}
          />,
        )}
      </Form.Item>
      <Form.Item>
        {getFieldDecorator('password', {
          rules: [{required: true, message: i18n.t('auth.password_required')}],
        })(
          <Input.Password
            prefix={<Icon type="lock" style={{color: 'rgba(0,0,0,.25)'}} />}
            placeholder={i18n.t('auth.password')}
          />,
        )}
      </Form.Item>
      <Form.Item>
        {getFieldDecorator('remember', {
          valuePropName: 'checked',
          initialValue: true,
        })(<Checkbox>{i18n.t('auth.remember')}</Checkbox>)}
      </Form.Item>
      <Button type="primary" htmlType="submit" className="login_form-button">
        {i18n.t('auth.login')}
      </Button>
      <div className="create_new">{i18n.t('auth.create_new')}</div>
    </Form>
  );
};

LoginForm.propTypes = {
  form: PropTypes.object,
  className: PropTypes.string,
  styles: PropTypes.object,
  submit: PropTypes.func,
};

LoginForm.defaultProps = {
  form: {},
  className: '',
  styles: {},
  submit: () => {
  },
};

export default Form.create()(LoginForm);
