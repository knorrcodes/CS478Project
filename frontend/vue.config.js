module.exports = {
  devServer: {
    host: "localhost",
    proxy: {
      "^/graphql": {
        target: "http://127.0.0.1:8080",
        changeOrigin: true
      }
    }
  }
};
