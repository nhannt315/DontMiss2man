import {CHANGE_LANGUAGE} from './actionTypes';

export const changeLanguage = lng => {
  return {
    type: CHANGE_LANGUAGE,
    payload: {language: lng},
  };
};
