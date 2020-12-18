# 得物面试

## 一面问题

### 简述 HTTPS 的加密与认证过程

> https(全称Hyper Text Transfer Protocol over Secure Socket Layer 超文本传输安全协议)

https在传统的http和tcp之间加了一层用于加解密的SSL/TLS层(安全套阶层 Secure Socket Layer/ 安全传输层 Transport Layer Security)。 使用https必须要有一套自己的数字证书(包含公钥和私钥)

https解决的问题：

- 信息加密传输：第三方无法窃听
- 校验机制：一旦被篡改，通信双方会立即发现
- 身份证书：防止身份被冒充

https加密过程：

- 1.客户端请求服务器获取证书公钥
- 2.客户端(ssl/tls)解析证书，校验证书是否有效
- 3.有效生成随机值
- 4.用公钥加密随机值生成秘钥
- 5.客户端将秘钥发送给服务器
- 6.服务端用私钥解密秘钥得到随机值
- 7.将信息和随机值混合在一起进行对称加密
- 8.将加密的内容发送给客户端
- 9.客户端用秘钥解密信息
