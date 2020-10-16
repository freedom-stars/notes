# 解决git clone 速度过慢的问题

git clone特别慢可能是因为github的域名被限制了,
通过代理
通过域名对应的ip地址，然后在hosts文件中加上ip–>域名的映射，刷新DNS缓存便可

## 1、代理

### http全局代理所有 不建议

```shell
git config --global http.proxy http://127.0.0.1:1087
git config --global https.proxy https://127.0.0.1:1087
```

### http 代理 github

```shell
git config --global http.https://github.com.proxy https://127.0.0.1:1087
git config --global https.https://github.com.proxy https://127.0.0.1:1087
```

### 使用socket代理 github ps: 仅代理github

```shell
git config --global http.https://github.com.proxy socks5://127.0.0.1:1086
git config --global https.https://github.com.proxy socks5://127.0.0.1:1086
```

### 设置post buff 建议都设置

```shell
git config --global http.postBuffer 524288000
```

## 2、域名映射

在网站 `https://www.ipaddress.com`解析地址

```shell
github.global.ssl.fastly.net
github.com
```

Windows上的hosts文件路径在C:\Windows\System32\drivers\etc\hosts
Mac/Linux的hosts文件路径在：sudo vim /etc/hosts

```shell
Vim /etc/hosts

# 解析出来的地址

199.232.69.194 github.global.ssl.fastly.net
140.82.113.4 github.com

```

最后刷新DNS即可

Windows刷新dns:

```shell
ipconfig /flushdns
```

Linux刷新dns:

```shell
systemctl restart nscd 或者 /etc/init.d/nscd restart
```

Mac刷新dns:

```shell
sudo killall -HUP mDNSResponder
```

注:如果速度还是很慢, 可在 `https://www.ipaddress.com` 多解析几次IP, 找到速度最快的ip地址放到host文件里映射即可
