import React from 'react';
import { useRouter } from 'next/router';
import useTranslation from 'next-translate/useTranslation';
import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import * as yup from 'yup';
import Warning from 'src/assets/svg/Warning.svg';
import AuthService from 'src/services/api/auth';
import { setAccessToken } from 'src/utils/cookie';
import { useAuth } from 'src/hooks/auth';

const RegistrationForm: React.FC = () => {
  const router = useRouter();
  const { setEmail, setToken } = useAuth();
  const { t } = useTranslation('auth');
  const { register, handleSubmit, formState, setError } = useForm<{
    email: string;
    password: string;
    passwordConfirmation: string;
  }>({
    defaultValues: {
      email: '',
      password: '',
      passwordConfirmation: '',
    },
    resolver: yupResolver(
      yup.object({
        email: yup.string().email().required(),
        password: yup.string().required(),
        passwordConfirmation: yup
          .string()
          .oneOf([yup.ref('password'), null], t('password_not_match')),
      })
    ),
  });

  const onSubmit = handleSubmit(async (data) => {
    try {
      const response = await AuthService.register(
        data.email,
        data.password,
        data.passwordConfirmation
      );
      setAccessToken(response.data.token);
      setEmail(response.data.email);
      setToken(response.data.token);
      router.push('/');
    } catch (error) {
      console.log(error);
    }
  });

  return (
    <div className="flex flex-col bg-white py-4 px-6 justify-center items-center h-96">
      <div className="flex flex-row items-center justify-center space-x-2">
        <img src="/logo.png" alt="Logo" className="w-10 h-10" />
        <span className="tracking-widest leading-relaxed uppercase font-bold text-2xl text-blue-400">
          DM2M
        </span>
      </div>
      <div className="w-full pt-4">
        <form onSubmit={onSubmit}>
          <div className="w-full">
            <label className="text-sm pb-1 text-gray-500" htmlFor="email">
              {t('email')}
            </label>
            <input
              id="email"
              {...register('email', { required: true })}
              className="p-2 placeholder-gray-300 shadow rounded-sm w-full"
              type="email"
              required
              placeholder="Email"
              autoComplete="email"
            />
          </div>
          <div className="w-full mt-4">
            <label className="text-sm pb-1 text-gray-500" htmlFor="password">
              {t('password')}
            </label>
            <input
              id="password"
              {...register('password', { required: true })}
              className="p-2 placeholder-gray-300 shadow rounded-sm w-full"
              type="password"
              required
              placeholder="Password"
              autoComplete="current-password"
            />
          </div>
          <div className="w-full mt-4">
            <label
              className="text-sm pb-1 text-gray-500"
              htmlFor="passwordConfirmation"
            >
              {t('password_confirm')}
            </label>
            <input
              id="passwordConfirmation"
              {...register('passwordConfirmation', { required: true })}
              className="p-2 placeholder-gray-300 shadow rounded-sm w-full"
              type="password"
              required
              placeholder="Password"
              autoComplete="current-password"
            />
          </div>
          {formState.isDirty && formState.errors.passwordConfirmation && (
            <div className="mt-3 flex text-red-500">
              <div className="flex justify-center">
                <Warning className="w-4 h-4" />
              </div>
              <div className="ml-2 text-xs">
                {formState.errors.passwordConfirmation.message}
              </div>
            </div>
          )}
          <div className="mt-10 flex flex-col justify-center">
            <button
              type="submit"
              className="login-button py-1 px-7 rounded-sm text-white bg-blue-400"
              disabled={formState.isSubmitting}
            >
              {t('register')}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default RegistrationForm;
