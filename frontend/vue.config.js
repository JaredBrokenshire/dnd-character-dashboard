const path = require('path')

module.exports = {
  publicPath: "/",
  devServer: {
    host: 'localdnd.com',
    liveReload: true,
  },
  configureWebpack: {
    resolve: {
      alias: {
        '@': path.resolve(__dirname, 'src'),
        '@core': path.resolve(__dirname, 'src/@core'),
      },
    },
  },
  transpileDependencies: true,
}
