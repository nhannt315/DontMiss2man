import React from 'react';
import RegistrationForm from 'src/components/registration/Form';

const Registration: React.FC = () => {
  return (
    <div
      className="w-full h-full flex flex-row items-center justify-center bg-grey"
      style={{
        backgroundImage:
          'linear-gradient(to right bottom, rgb(99, 125, 143), rgba(52, 52, 52, 0.6)), url(bg.jpg)',
      }}
    >
      <div className="w-4/12">
        <RegistrationForm />
      </div>
    </div>
  );
};

export default Registration;
