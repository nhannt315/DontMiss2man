import React from 'react';
import useTranslation from 'next-translate/useTranslation';
import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import * as yup from 'yup';
import Warning from 'src/assets/svg/Warning.svg';
import axios from 'axios';
import { configureAxios } from 'src/config/axios';

import AuthService from 'src/services/api/auth';

const LoginForm: React.FC = () => {
  const { t } = useTranslation('auth');
  const { register, handleSubmit, formState, setError } = useForm<{
    email: string;
    password: string;
  }>({
    defaultValues: {
      email: '',
      password: '',
    },
    resolver: yupResolver(
      yup.object().shape({
        email: yup.string().email().required(),
        password: yup.string().required(),
      })
    ),
  });

  const onSubmit = handleSubmit(async (data) => {
    configureAxios('localhost:4000/api/v1');
    console.log(axios.defaults.baseURL);

    AuthService.login(data.email, data.password)
      .then((response) => {
        console.log(response.data);
      })
      .catch((e) => {
        console.log(e);
      });
    setError('password', { message: '認証に失敗しました' });
  });

  return (
    <div className="flex flex-col bg-white py-4 px-6 justify-center items-center h-80">
      <div className="flex flex-row items-center justify-center space-x-2">
        <img src="/logo.png" alt="Logo" className="w-10 h-10" />
        <span className="tracking-widest leading-relaxed uppercase font-bold text-2xl text-blue-400">
          DM2M
        </span>
      </div>
      <div className="w-full pt-4">
        <form onSubmit={onSubmit}>
          <div className="w-full">
            <input
              {...register('email', { required: true })}
              className="p-2 placeholder-gray-300 shadow rounded-sm w-full"
              type="email"
              required
              placeholder="Email"
              autoComplete="email"
            />
          </div>
          <div className="w-full mt-4">
            <input
              {...register('password', { required: true })}
              className="p-2 placeholder-gray-300 shadow rounded-sm w-full"
              type="password"
              required
              placeholder="Password"
              autoComplete="current-password"
            />
          </div>
          {formState.isDirty && formState.errors.password && (
            <div className="mt-3 flex text-gray-500">
              <div className="flex justify-center">
                <Warning className="w-4 h-4" />
              </div>
              <div className="ml-2 text-xs">
                {formState.errors.password.message}
              </div>
            </div>
          )}
          <div className="mt-10 flex justify-center">
            <button
              type="submit"
              className="login-button py-1 px-7 rounded-sm text-white bg-blue-400"
              disabled={formState.isSubmitting}
            >
              {t('login')}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default LoginForm;
