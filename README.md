

# Gredir

一个简洁的支持多用户多端口tcp转发工具，该工具支持作为大部分矿池的中转代理服务器。同时支持tls加密选项。该工具使用golang开发。

<!-- PROJECT SHIELDS -->

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

## 目录

- [安装](#部署)
- [配置说明](#配置说明)
- [作者](#作者)

### 安装

```shell
tar xzvf gredir*.tar.gz && cd gredir && bash install.sh
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
[your-project-path]:shaojintian/Best_README_template
[contributors-shield]: https://img.shields.io/github/contributors/shaojintian/Best_README_template.svg?style=flat-square
[contributors-url]: https://github.com/shaojintian/Best_README_template/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/shaojintian/Best_README_template.svg?style=flat-square
[forks-url]: https://github.com/shaojintian/Best_README_template/network/members
[stars-shield]: https://img.shields.io/github/stars/shaojintian/Best_README_template.svg?style=flat-square
[stars-url]: https://github.com/shaojintian/Best_README_template/stargazers
[issues-shield]: https://img.shields.io/github/issues/shaojintian/Best_README_template.svg?style=flat-square
[issues-url]: https://img.shields.io/github/issues/shaojintian/Best_README_template.svg
[license-shield]: https://img.shields.io/github/license/shaojintian/Best_README_template.svg?style=flat-square
[license-url]: https://github.com/shaojintian/Best_README_template/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/shaojintian
