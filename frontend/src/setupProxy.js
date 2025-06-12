const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = function(app) {
  // 代理用户相关请求
  app.use(
    '/user',
    createProxyMiddleware({
      target: 'http://localhost:8888',
      changeOrigin: true,
    })
  );

  // 代理车辆相关请求
  app.use(
    '/vehicle',
    createProxyMiddleware({
      target: 'http://localhost:8888',
      changeOrigin: true,
    })
  );

  // 代理车辆类型相关请求
  app.use(
    '/vehicle-type',
    createProxyMiddleware({
      target: 'http://localhost:8888',
      changeOrigin: true,
    })
  );

  // 代理商户相关请求
  app.use(
    '/merchant',
    createProxyMiddleware({
      target: 'http://localhost:8888',
      changeOrigin: true,
    })
  );

  // 代理管理员相关请求
  app.use(
    '/admin',
    createProxyMiddleware({
      target: 'http://localhost:8888',
      changeOrigin: true,
    })
  );
};