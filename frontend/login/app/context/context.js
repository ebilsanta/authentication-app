'use client'
import { createContext, useContext, useState } from 'react';

const LoginContext = createContext();

export const LoginProvider = ({ children }) => {
  const [hostedLogin, setHostedLogin] = useState(false);

  const setHostedLoginContext = (value) => {
    setHostedLogin(value);
  };

  return (
    <LoginContext.Provider value={{ hostedLogin, setHostedLoginContext }}>
      {children}
    </LoginContext.Provider>
  );
};

export const useLogin = () => {
  return useContext(LoginContext);
};