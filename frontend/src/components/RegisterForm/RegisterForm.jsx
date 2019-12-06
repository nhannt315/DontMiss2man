import React from 'react';
import PropTypes from 'prop-types';
import {Form, Input, Button, Icon} from 'antd';

const RegisterForm = ({form, submit}) => {

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
    <Form onSubmit={handleSubmit}>
      <Form.Item>
        
      </Form.Item>
    </Form>
  );
};

RegisterForm.propTypes = {
  form: PropTypes.object,
  submit: PropTypes.func,
};

RegisterForm.defaultProps = {
  form: {},
  submit: () => {
  },
};

export default Form.create()(RegisterForm);
