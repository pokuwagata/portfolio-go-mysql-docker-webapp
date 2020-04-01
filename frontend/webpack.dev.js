const path = require('path');
const merge = require('webpack-merge');
const common = require('./webpack.common.js');

module.exports = merge(common, {
  mode: 'development',
  devtool: 'inline-source-map',
  devServer: {
    contentBase: path.join(__dirname, 'dist'),
    compress: true,
    port: 8080,
    historyApiFallback: true,
    proxy: {
      '/api': {
        target: 'http://api:5000',
        pathRewrite: {'^/api' : ''}
      },
      '/bff': {
        target: 'http://bff:4000',
        pathRewrite: {'^/bff' : ''}
      }
    }
  }
});