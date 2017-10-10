const path = require('path');
// const webpack = require('webpack');
const Merge = require('webpack-merge');
const BaseConfig = require('./webpack.base.config');
const HtmlWebpackPlugin = require('html-webpack-plugin');

const PORT = 8500;

const HtmlWebpackPluginConfig = new HtmlWebpackPlugin({
  template: './src/index.html',
  filename: 'index.html',
  inject: 'body',
  title: 'HackGT Sponsorship Portal',
});

module.exports = Merge(BaseConfig, {
  output: {
    path: path.resolve('../backend/public'),
    filename: 'bundle.js',
  },
  devtool: '#source-map',
  plugins: [HtmlWebpackPluginConfig],
  devServer: {
    port: PORT,
    stats: 'minimal',
  },
});
