import React, { ReactElement, useEffect, useState } from 'react';
import App from 'next/app';
import type { AppProps, AppContext } from 'next/app';
import Head from 'next/head';
import { configureAxios } from 'src/config/axios';
import { BASE_API_URL, SSR_API_URL } from 'src/constants/endpoint';
import { AuthProvider } from 'src/hooks/auth';
import Layout from 'src/components/Layout';
import UserService from 'src/services/api/user';
import { getAccessToken } from 'src/utils/cookie';
import { IUserInfo } from 'src/types/user';
import 'src/style.css';

interface IAppProps extends AppProps {
  baseApiUrl: string;
  accessToken: string;
}

const MyApp = ({
  Component,
  pageProps,
  baseApiUrl,
  accessToken,
}: IAppProps): ReactElement => {
  configureAxios(baseApiUrl || BASE_API_URL);
  const [userInfo, setUserInfo] = useState<IUserInfo | null>(null);
  useEffect(() => {
    if (!accessToken) return;
    UserService.getUserInfo(accessToken)
      .then((res) => {
        setUserInfo({ token: accessToken, email: res.data.email });
      })
      .catch((err) => {
        console.log(err);
      });
  }, [accessToken]);

  return (
    <AuthProvider>
      <Layout userInfo={userInfo}>
        <Head>
          <title>DontMiss2Man Renewal</title>
          <meta name="viewport" content="width=device-width, initial-scale=1" />
        </Head>
        <Component {...pageProps} />
      </Layout>
    </AuthProvider>
  );
};

MyApp.getInitialProps = async (appContext: AppContext) => {
  // calls page's `getInitialProps` and fills `appProps.pageProps`
  const appProps = await App.getInitialProps(appContext);

  let baseUrl: string;
  // Initialize axios
  if (appContext.ctx.req) {
    // Sever side
    baseUrl = SSR_API_URL;
  } else {
    // Client side
    baseUrl = BASE_API_URL;
  }

  const accessToken = getAccessToken(appContext.ctx);
  return { ...appProps, baseApiUrl: baseUrl, accessToken: accessToken };
};

export default MyApp;
