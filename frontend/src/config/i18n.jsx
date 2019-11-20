import i18next from 'i18next';
import lang from './locales/lang';

i18next.init({
  interpolation: {
    escapeValue: false,
  },
  lng: 'ja',
  resources: lang,
});
export default i18next;
