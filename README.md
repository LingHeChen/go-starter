# Go Starter / Go Starter 项目模板

![License](https://img.shields.io/badge/license-MIT-blue.svg) ![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue) ![SvelteKit](https://img.shields.io/badge/SvelteKit-Latest-orange)

[English](#english) | [中文](#chinese)

<a name="english"></a>
## 🇬🇧 English

**Go Starter** is a modern full-stack project template designed to help you hit the ground running. It combines a robust **Go** backend with a high-performance **SvelteKit** frontend, all managed easily via a `Makefile`.

### ✨ Features

-   **Full-Stack**: Integrated Go backend and SvelteKit frontend.
-   **Developer Experience**: One-command live reload for both backend (`air`) and frontend (`vite`).
-   **Self-Contained**: Local tool installation (Air, Swag) to avoid global dependency conflicts.
-   **Modular**: Structured with clean architecture principles.

### 🛠 Prerequisites

Ensure you have the following installed:
-   [Go](https://go.dev/) (1.21+)
-   [Node.js](https://nodejs.org/) & [pnpm](https://pnpm.io/)
-   Make

### 🚀 Getting Started

1.  **Initialize Project**
    Installs necessary Go tools (Air, Swag) locally and JavaScript dependencies.
    ```bash
    make init
    ```

2.  **Start Development**
    Runs both backend and frontend in watch mode.
    ```bash
    make dev
    ```
    -   Backend runs on: `http://localhost:8080` (or configured port)
    -   Frontend runs on: `http://localhost:5173`

### 📦 Build

Build both the frontend and backend.

```bash
make build
```

### 📚 Documentation

Generate Swagger API documentation:

```bash
make docs
```

---

<a name="chinese"></a>
## 🇨🇳 中文

**Go Starter** 是一个现代化的全栈项目模板，旨在帮助您快速启动开发。它结合了强大的 **Go** 后端和高性能的 **SvelteKit** 前端，并通过 `Makefile` 进行统一管理。

### ✨ 特性

-   **全栈开发**: 集成 Go 后端与 SvelteKit 前端。
-   **极致体验**: 一键启动前后端热重载开发环境（集成 `air` 和 `vite`）。
-   **工具隔离**: 自动化安装本地开发工具（Air, Swag），避免污染全局环境。
-   **模块化**: 遵循清晰的架构原则设计。

### 🛠 前置要求

请确保您的环境已安装：
-   [Go](https://go.dev/) (1.21+)
-   [Node.js](https://nodejs.org/) 和 [pnpm](https://pnpm.io/)
-   Make 工具

### 🚀 快速开始

1.  **初始化项目**
    安装必要的 Go 工具（Air, Swag）到本地目录，并安装前端依赖。
    ```bash
    make init
    ```

2.  **启动开发环境**
    同时启动后端和前端的监听模式。
    ```bash
    make dev
    ```
    -   后端地址: `http://localhost:8080` (默认)
    -   前端地址: `http://localhost:5173`

### 📦 构建

构建前端静态资源并编译 Go 二进制文件。

```bash
make build
```

### 📚 文档

生成 Swagger API 文档：

```bash
make docs
```
