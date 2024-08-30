# BetaX Blog - 轻量级动态博客平台

[English](README.md)

[![License](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

## 介绍

BetaX Blog 是一个用 Go 语言编写的开源动态博客平台, 其轻量级的设计使其非常适合个人博客或小型社区使用

整个应用被封装成一个仅 17MB 大小的单一二进制文件，确保了快速部署和最小资源占用

同时使用接口数据缓存与PWA技术(需配置HTTPS)提高访问速度降低资源开销

但 BetaX Blog 终归还是一个动态博客, 如需更高的性能和效率请搭配CDN部署

## 特性

- **📄 Markdown 支持**: 支持 Markdown 撰写文章，便于格式化
- **🌕 暗黑模式**: 内置暗黑模式，适合低光环境下阅读
- **⏫ 图片上传**: 允许直接上传图片至文章中
- **🧠 AI 摘要**: 提供利用 AI 生成文章摘要功能
- **🔗 GitHub 集成**: 通过 GitHub 授权登录，并自动绑定博主信息
- **☁️ WebDAV 同步**: 支持通过 WebDAV 每 4 小时同步一次数据库于资源文件
- **⏬️ PWA 支持**: 通过 HTTPS 访问时 PWA 能给访问者更好的体验, 并有助于降低请求数量

## 开始使用

### 安装

#### 环境要求

- **Linux**: 目前仅支持 Linux 系统
- **curl**: 用于下载安装包

#### 快速启动

在终端执行以下命令：

```shell
bash -c "$(curl -fsSL https://betax.dev/sc/blog.sh)"
```

#### 离线安装

请在可连网的设备上下载相应的安装包，参见 [releases](https://github.com/skye-z/betax-blog/releases)

然后下载这个 [脚本](https://betax.dev/sc/blog.sh)，并将脚本和安装包一起上传到服务器上的同一目录下

运行脚本并选择 `Install BetaX Blog (Offline)`

#### 从源码构建

1. 克隆此仓库：
   ```bash
   git clone https://github.com/skye-z/betax-blog.git
   cd betax-blog
   ```
2. 构建应用程序：
   ```bash
   CGO_ENABLED=0 go build -o betax-blog -ldflags '-s -w'
   ```
3. 启动服务：
   ```bash
   ./betax-blog
   ```

### 配置

原则上不建议直接修改配置文件，因为修改后需要重启服务才能生效

建议在后台管理界面进行配置修改

### 使用说明

- 访问 `http://localhost:9800` (启动时使用 `--port` 参数指定端口) 以访问博客
- 完成初始化配置
- 绑定 Github 账户

### 贡献指南

欢迎您的贡献！请遵循以下步骤:

1. fork 项目
2. 创建分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 提交拉取请求 (pull request)

### 许可

本项目采用 GNU 通用公共许可协议第 3 版 (GNU General Public License v3.0) 许可 - 详情请参见 [LICENSE](LICENSE) 文件.
