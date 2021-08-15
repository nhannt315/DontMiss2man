const CracoAntDesignPlugin = require('craco-antd');
const emotion = require('babel-plugin-emotion');

process.env.BROWSER = 'none';

module.exports = {
  plugins: [
    {
      plugin: emotion
    },
    {
      plugin: CracoAntDesignPlugin
    }
  ]
};
