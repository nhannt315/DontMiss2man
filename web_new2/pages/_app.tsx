import React, { ReactElement } from 'react';
import App from 'next/app';
import type { AppProps, AppContext } from 'next/app';
import Head from 'next/head';
import { configureAxios } from 'src/config/axios';
import { BASE_API_URL, SSR_API_URL } from 'src/constants/endpoint';
import Layout from 'src/components/Layout';
import 'src/style.css';

const MyApp = ({ Component, pageProps }: AppProps): ReactElement => {
  return (
    <Layout>
      <Head>
        <title>DontMiss2Man Renewal</title>
        <meta name="viewport" content="width=device-width, initial-scale=1" />
      </Head>
      <Component {...pageProps} />
    </Layout>
  );
};

MyApp.getInitialProps = async (appContext: AppContext) => {
  // calls page's `getInitialProps` and fills `appProps.pageProps`
  const appProps = await App.getInitialProps(appContext);

  // Initialize axios
  if (appContext.ctx.req) {
    // Sever side
    configureAxios(SSR_API_URL);
  } else {
    // Client side
    configureAxios(BASE_API_URL);
  }

  console.log('I`m here', SSR_API_URL);

  return { ...appProps };
};

export default MyApp;
