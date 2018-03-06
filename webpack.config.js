var webpack = require("webpack");
var path = require('path');

module.exports = {
  entry:{ 
    app: './src/js/app.js',
  },
  output: {
    path: path.resolve(__dirname, './public/js'),
    filename: '[name].bundle.min.js'
   }
  // plugins: [
  //   new webpack.ProvidePlugin({
  //     $: 'jquery',
  //     'jQuery': 'jquery'
  //   })
  // ]
};