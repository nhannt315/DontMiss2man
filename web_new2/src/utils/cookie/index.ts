import { CookieSerializeOptions } from 'cookie';
import { NextPageContext } from 'next';
import nookies from 'nookies';

const COOKIE_NAMES = {
  ACCESS_TOKEN: 'access_token',
  LOCALE: 'NEXT_LOCALE',
} as const;

const DEFAULT_COOKIE_SERIALIZE_OPTIONS: CookieSerializeOptions = {
  secure: true,
  sameSite: 'lax',
} as const;

export const getAccessToken = (
  ctx?: Pick<NextPageContext, 'req'>
): string | undefined => {
  return nookies.get(ctx)[COOKIE_NAMES.ACCESS_TOKEN];
};
export const setAccessToken = (
  token: string,
  ctx?: Pick<NextPageContext, 'res'>
): void => {
  nookies.set(
    ctx,
    COOKIE_NAMES.ACCESS_TOKEN,
    token,
    DEFAULT_COOKIE_SERIALIZE_OPTIONS
  );
};
export const removeAccessToken = (ctx?: Pick<NextPageContext, 'res'>): void => {
  nookies.destroy(ctx, COOKIE_NAMES.ACCESS_TOKEN);
};

export const getLocale = (
  ctx?: Pick<NextPageContext, 'req'>
): string | undefined => {
  return nookies.get(ctx)[COOKIE_NAMES.LOCALE];
};

export const setLocale = (
  token: string,
  ctx?: Pick<NextPageContext, 'res'>
): void => {
  nookies.set(
    ctx,
    COOKIE_NAMES.LOCALE,
    token,
    DEFAULT_COOKIE_SERIALIZE_OPTIONS
  );
};
