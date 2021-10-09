import React, { useState, useEffect } from 'react';
import Link from 'next/link';
import Header from 'src/components/Header';
import { IUserInfo } from 'src/types/user';
import { useAuth } from 'src/hooks/auth';

interface IProps {
  userInfo?: IUserInfo;
}

const MainLayout: React.FC<IProps> = ({ userInfo, children }) => {
  const [showHeader, setShowHeader] = useState<boolean>(true);
  const { setEmail, setToken } = useAuth();

  useEffect(() => {
    if (userInfo) {
      setEmail(userInfo.email);
      setToken(userInfo.token);
    }
  }, [setEmail, setToken, userInfo]);

  return (
    <div id="app-bar" style={{ height: '100vh' }}>
      {/*Modal goes here*/}
      <div className="flex flex-col bg-gray-200 h-full">
        <Header show={showHeader} />
        <div className="flex-1">{children}</div>
      </div>
    </div>
  );
};

export default MainLayout;
