import React, { useState } from 'react';
import Link from 'next/link';
import Header from 'src/components/Header';

const MainLayout: React.FC = ({ children }) => {
  const [showHeader, setShowHeader] = useState<boolean>(true);

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
