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