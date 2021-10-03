import React, { createContext, useContext, useState, ReactNode } from 'react';

interface IAuthInfo {
  token: string;
  email: string;
  setToken: (token: string) => void;
  setEmail: (email: string) => void;
}

const authInfoDefaulValue: IAuthInfo = {
  token: null,
  email: null,
  setToken: (token: string) => token,
  setEmail: (email: string) => email,
};

type Props = {
  children: ReactNode;
};

export const AuthContext = createContext<IAuthInfo>(authInfoDefaulValue);
export const useAuth = () => useContext(AuthContext);

export const AuthProvider = ({ children }: Props) => {
  const [token, setToken] = useState<string>(null);
  const [email, setEmail] = useState<string>('');

  const value: IAuthInfo = {
    token,
    email,
    setToken,
    setEmail,
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};
