const path = require('path');

module.exports = {
  stories: ['../src/**/*.stories.@(tsx|jsx)'],
  addons: [
    '@storybook/addon-actions',
    '@storybook/addon-links',
    '@storybook/addon-controls',
    // PostCSS 8+のときの設定
    // ref: https://storybook.js.org/addons/@storybook/addon-postcss
    {
      name: '@storybook/addon-postcss',
      options: {
        postcssLoaderOptions: {
          implementation: require('postcss'),
        },
      },
    },
  ],
  typescript: {
    // TODO: ビルドエラーによる一時的な対策
    // ref: https://github.com/storybookjs/storybook/issues/15067
    // ref: https://github.com/styleguidist/react-docgen-typescript/issues/356
    reactDocgen: 'none',
  },
  webpackFinal: async (config, env) => {
    // svgの読み込みにデフォルトでfile-loaderが当たってるので、置き換える
    const fileLoaderRule = config.module.rules.find(
      (rule) => rule.test && rule.test.test('.svg')
    );
    fileLoaderRule.exclude = /\.svg$/;
    config.module.rules.push({
      test: /\.svg$/,
      use: ['@svgr/webpack'],
    });

    // importパスの起点を追加
    config.resolve.modules = [
      ...(config.resolve.modules || []),
      path.resolve(__dirname, '../'),
      path.resolve(__dirname, '../src'),
    ];

    return config;
  },
};
