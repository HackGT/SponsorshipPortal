const path = require('path');
const ExtractTextPlugin = require('extract-text-webpack-plugin');

const ExtractTextPluginConfig = new ExtractTextPlugin({
  filename: 'bundle.css',
  disable: false,
  allChunks: true,
});

const BabelLoaderConfig = {
  loader: 'babel-loader',
  query: {
    presets: ['es2015', 'stage-2', 'react'],
    plugins: ['react-html-attrs'],
  },
};

module.exports = {
  entry: ['./client/js/index.jsx'],
  module: {
    rules: [
      {
        test: /\.js$/,
        use: [BabelLoaderConfig, 'eslint-loader'],
        include: path.resolve(__dirname, 'client'),
      },
      {
        test: /\.jsx$/,
        use: [BabelLoaderConfig, 'eslint-loader'],
        include: path.resolve(__dirname, 'client'),
      },
      { test: /\.css$/,
        use: ExtractTextPlugin.extract({
          fallback: 'style-loader',
          use: [
            'css-loader',
          ],
        }),
      },
      { test: /\.(scss|sass)$/,
        use: ExtractTextPlugin.extract({
          fallback: 'style-loader',
          use: [
            'css-loader',
            'postcss-loader',
            'sass-loader',
          ],
        }),
      },
      {
        test: /\.(jpe?g|png|gif|svg)$/i,
        use: [
          'file-loader?name=[path][name].[ext]',
        ],
      },
      {
        test: /\.(eot|ttf|woff|woff2)$/i,
        use: ['file-loader?[name].[ext]&mimetype=application/x-font-truetype'],
      },
    ],
    // loaders: [
    //   {
    //     test: /particles\.js/,
    //     loader: 'exports?particlesJS=window.particlesJS,pJSDom=window.pJSDom',
    //   },
    // ],
  },
  plugins: [ExtractTextPluginConfig],
  resolve: {
    extensions: ['.js', '.jsx'],
  },
};
