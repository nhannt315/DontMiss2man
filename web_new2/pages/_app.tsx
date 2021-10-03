import React, { ReactElement } from 'react';
import App from 'next/app';
import type { AppProps, AppContext } from 'next/app';
import Head from 'next/head';
import { configureAxios } from 'src/config/axios';
import { BASE_API_URL, SSR_API_URL } from 'src/constants/endpoint';
import { AuthProvider } from 'src/hooks/auth';
import Layout from 'src/components/Layout';
import 'src/style.css';

interface IAppProps extends AppProps {
  baseApiUrl: string;
}

const MyApp = ({
  Component,
  pageProps,
  baseApiUrl,
}: IAppProps): ReactElement => {
  configureAxios(baseApiUrl || BASE_API_URL);
  return (
    <Layout>
      <Head>
        <title>DontMiss2Man Renewal</title>
        <meta name="viewport" content="width=device-width, initial-scale=1" />
      </Head>
      <AuthProvider>
        <Component {...pageProps} />
      </AuthProvider>
    </Layout>
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

  return { ...appProps, baseApiUrl: baseUrl };
};

export default MyApp;
