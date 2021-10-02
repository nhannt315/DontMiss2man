import i18next from 'i18next';
import intervalPlural from 'i18next-intervalplural-postprocessor';
import lang from './locales/lang';

i18next
  .use(intervalPlural)
  .init({
  interpolation: {
    escapeValue: false,
  },
  lng: 'ja',
  resources: lang,
});
export default i18next;
