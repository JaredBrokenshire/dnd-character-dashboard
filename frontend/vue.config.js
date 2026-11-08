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
        '@core': path.resolve(__dirname, 'src/@core'),
        '@axios': path.resolve(__dirname, 'src/services/api'),
      },
    },
  },
  transpileDependencies: true,
}
