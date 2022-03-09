# Gredir

一个支持多用户多端口tcp转发工具，该工具支持作为大部分矿池的中转代理服务器。同时支持tls加密选项。该工具使用golang开发。

<!-- PROJECT SHIELDS -->

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![GPL License][license-shield]][license-url]

## 目录

- [安装](#部署)
- [配置说明](#配置说明)
- [作者](#作者)

### 安装

```shell
tar xzvf gredir*.tar.gz && cd gredir && bash install.sh
```

### 生成自签名证书

```shell
# 创建根证书私钥，长度2048
openssl genrsa -out ca.key 2048
# 创建根证书
openssl req -new -x509 -days 36500 -key ca.key -out ca.crt
# 创建ssl证书私钥
openssl genrsa -out server.key 2048
# 创建ssl证书，xxxx改为奇数
openssl req -new -key server.key -out server.csr
mkdir dir demoCA &&cd demoCA&&mkdir newcerts&&echo 'xxxx' > serial &&touch index.txt&&cd ..
# 用CA根证书签署ssl证书，生成server.crt与server.key
openssl ca -in server.csr -out server.crt -cert ca.crt -keyfile ca.key
```

### 配置说明

```yaml
proxies:
  - id: 0
    local_addr: ":20000" //中转服务器转发端口
    remote_addr: "btc.f2pool.com:1314" //待转发的远程目的服务器
    user: "zcy" 
    date: "2022-2-22"
    tls: true //是否启用tls转发
    
log: "/var/log/gredir.log"  //log文件路径
  
tls:
  public_file: "/usr/local/gredir/server.crt" //tls证书路径
  private_file: "/usr/local/gredir/server.key" //密钥路径
```

### 作者

zcylove1995@gmail.com

### 版权说明

该项目签署了MIT 授权许可，详情请参阅 [LICENSE.txt](https://github.com/shaojintian/Best_README_template/blob/master/LICENSE.txt)

<!-- links -->
[your-project-path]:wytfy/gredir
[contributors-shield]: https://img.shields.io/github/contributors/wytfy/gredir.svg?style=flat-square
[contributors-url]: https://github.com/wytfy/gredir/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/wytfy/gredir.svg?style=flat-square
[forks-url]: https://github.com/wytfy/gredir/network/members
[stars-shield]: https://img.shields.io/github/stars/wytfy/gredir.svg?style=flat-square
[stars-url]: https://github.com/wytfy/gredir/stargazers
[issues-shield]: https://img.shields.io/github/issues/wytfy/gredir.svg?style=flat-square
[issues-url]: https://img.shields.io/github/issues/wytfy/gredir.svg
[license-shield]: https://img.shields.io/github/license/wytfy/gredir.svg?style=flat-square
[license-url]: https://github.com/wytfy/gredir/blob/master/LICENSE.txt
