export type AccessToken = {
  accessToken: string;
  tokenType: string;
  uid: string;
  client: string;
};

export const generateRequestHeader = (token: AccessToken) => {
  return {
    'access-token': token.accessToken,
    'token-type': token.tokenType,
    uid: token.uid,
    client: token.client,
  };
};
