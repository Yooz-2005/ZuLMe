// 简单的JWT token生成脚本用于测试
const jwt = require('jsonwebtoken');

// 用户JWT密钥（从后端代码中获取）
const secretKey = '2209';

// 创建测试用户的claims
const claims = {
  ID: 1,  // 测试用户ID
  NickName: 'testuser',
  AuthorityId: 1,
  exp: Math.floor(Date.now() / 1000) + (60 * 60 * 24), // 24小时后过期
  iat: Math.floor(Date.now() / 1000)
};

// 生成token
const token = jwt.sign(claims, secretKey);

console.log('测试用户Token:');
console.log(token);
console.log('\n请在浏览器控制台中运行以下命令来设置token:');
console.log(`localStorage.setItem('token', '${token}');`);
console.log('\n然后刷新页面并尝试预订功能。');
