var webpack = require("webpack");
var path = require('path');

module.exports = {
  entry:{ 
    app: './wwwroot/app/app.js',
  },
  output: {
    path: path.resolve(__dirname, 'wwwroot/public/js'),
    filename: '[name].bundle.min.js'
   }
  // plugins: [
  //   new webpack.ProvidePlugin({
  //     $: 'jquery',
  //     'jQuery': 'jquery'
  //   })
  // ]
};