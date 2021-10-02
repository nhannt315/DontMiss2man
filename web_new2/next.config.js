const withPlugins = require("next-compose-plugins");
const nextTranslate = require("next-translate");

const nextConfig = {
    reactStrictMode: true,
    webpack: (config) => {
        // for svg import
        config.module.rules.push({
            test: /\.svg$/,
            use: ['@svgr/webpack'],
        });

        // for font file
        config.module.rules.push({
            test: /\.(eot|ttf|woff2|otf)$/,
            use: 'url-loader?limit=8192&name=images/[name].[ext]',
        });

        return config;
    },
};

module.exports = withPlugins([
    [nextTranslate]
], nextConfig);
