export const generateRequestHeader = (token: string) => {
  return {
    Authorization: `Bearer ${token}`,
  };
};
