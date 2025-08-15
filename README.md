# Go Hello World Web 应用

这是一个使用Go语言实现的简单Web应用，提供HTTP接口返回"Hello World"。

## 功能特性

- 简单的HTTP服务器
- 根路径返回"Hello World"消息
- 默认监听8080端口

## 运行方法

1. 确保已安装Go语言环境
2. 在项目根目录执行：
   ```bash
   go run main.go
   ```
3. 打开浏览器访问：http://localhost:8080

## 测试接口

可以使用curl命令测试：
```bash
curl http://localhost:8080
```

预期输出：`Hello World!`

## Docker 部署

### 构建镜像
```bash
docker build -t go-hello-world .
```

### 运行容器
```bash
docker run -p 8080:8080 go-hello-world
```

### 后台运行
```bash
docker run -d -p 8080:8080 --name hello-app go-hello-world
```

### 查看运行状态
```bash
docker ps
docker logs hello-app
```

## GitHub Actions 自动构建

本项目配置了GitHub Actions来自动构建和推送Docker镜像到GitHub Container Registry (ghcr.io)。

### 自动触发条件

1. **推送代码**：推送到 `main` 或 `master` 分支时自动构建
2. **创建标签**：创建 `v*` 格式的标签时自动构建
3. **发布版本**：在GitHub上发布Release时自动构建

### 镜像标签规则

- `latest`：最新版本
- `v1.0.0`：具体版本号
- `main-abc123`：分支名+提交哈希

### 使用ghcr.io镜像

```bash
# 拉取最新镜像
docker pull ghcr.io/你的用户名/go-hello-1:latest

# 拉取特定版本
docker pull ghcr.io/你的用户名/go-hello-1:v1.0.0

# 运行容器
docker run -p 8080:8080 ghcr.io/你的用户名/go-hello-1:latest
```

### 权限配置

确保仓库设置中启用了：
- Actions → General → Workflow permissions → "Read and write permissions"
- Packages → "Inherit access from source repository"