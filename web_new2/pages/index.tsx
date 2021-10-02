import React from 'react';
import { useRouter } from 'next/router';

const redirectTo = '/home';

const Index = () => {
  const router = useRouter();
  if (typeof window !== 'undefined') {
    router.push(redirectTo);
  }
  return <></>;
};

Index.getInitialProps = async ({ ctx }) => {
  if (ctx && ctx.req) {
    ctx.res.statusCode = 302;
    ctx.res.setHeader('Location', redirectTo);
  }
  return { props: '' };
};

export default Index;
