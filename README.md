# BetaX Blog - A Lightweight Dynamic Blogging

[中文](README_zh.md)

[![License](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

## Introduction

BetaX Blog is an open-source, dynamic blogging platform written in Go. Its lightweight design makes it ideal for personal blogs or small communities.

The entire application is packaged into a single binary that weighs only 17MB, ensuring quick deployment and minimal resource usage.

At the same time using interface data caching and PWA technology (need to configure HTTPS) to improve access speed and reduce resource overhead.

However, BetaX Blog is still a dynamic blog, for higher performance and efficiency, please deploy with CDN.

## Features

- **📄 Markdown Support**: Write your posts using Markdown for easy formatting.
- **🌕 Dark Mode**: Built-in dark mode for comfortable reading in low-light environments.
- **⏫ Image Upload**: Directly upload images to your posts.
- **🧠 AI Summarization**: Automatically generate summaries for your articles using AI.
- **🔗 GitHub Integration**: Authenticate users via GitHub for seamless login and user management.
- **☁️ WebDAV Synchronization**: Supports synchronization of database and resource files via WebDAV every 4 hours.
- **⏬️ PWA Support**: PWAs provide a better experience for visitors when accessing over HTTPS and help reduce the number of requests.

## Getting Started

### Installation

#### Requirements

- **linux**: Currently only supports Linux systems
- **curl**: Used for downloading installation packages

#### Quick Start

Execute the ⬇️ command in the terminal

```shell
bash -c "$(curl -fsSL https://betax.dev/sc/blog.sh)"
```

#### Offline installation

Please download the corresponding installation package in the [releases](https://github.com/skye-z/betax-blog/releases) on devices that can be connected to the internet.

Then download [this script](https://betax.dev/sc/blog.sh). Upload the script and installation package together to the same directory on the server.

Execute the script and select `Install BetaX Blog (Offline)`.

#### From Source

1. Clone this repository:
   ```bash
   git clone https://github.com/skye-z/betax-blog.git
   cd betax-blog
   ```
2. Build the application:
   ```bash
   CGO_ENABLED=0 go build -o betax-blog -ldflags '-s -w'
   ```
3. Start the server:
   ```bash
   ./betax-blog
   ```

### Configuration

In principle, it is not recommended to directly modify the configuration file, as the modification requires a restart of the service to take effect

Suggest modifying the configuration in the management backend.

### Usage

- Visit `http://localhost:9800` (use `--port` to specify the port) to access the blog.
- Complete initialization configuration
- Bind Github account

### Contributing

We welcome contributions! Please follow these guidelines:

1. Fork the project.
2. Create a feature branch (`git checkout -b feature/amazing-feature`).
3. Commit your changes (`git commit -m 'Add some amazing feature'`).
4. Push to the branch (`git push origin feature/amazing-feature`).
5. Open a pull request.

### License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.
