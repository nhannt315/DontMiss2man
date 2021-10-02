module.exports = {
  locales: ['en', 'ja', 'vi'],
  defaultLocale: 'ja',
  pages: {
    '*': [
      'common',
      'auth',
      'confirmation',
      'homepage',
      'language',
      'roomDetail',
      'searchFilter',
      'sideMenu',
    ],
  },
  loadLocaleFrom: (lang, ns) =>
    import(`./public/locales/${lang}/${ns}.json`).then((m) => m.default),
};
