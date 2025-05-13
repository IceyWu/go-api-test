# Test API

这是一个使用 Golang 和 Gin 框架构建的 RESTful API 服务器示例。

## 项目结构

```
.
├── main.go              # 主程序入口
├── internal/            # 内部包
│   ├── handler/        # 请求处理器
│   ├── model/          # 数据模型
│   └── router/         # 路由配置
└── docs/               # Swagger 文档
```

## 功能特性

- RESTful API 设计
- Swagger API 文档
- 模块化的项目结构
- 用户管理 API

## API 端点

### 用户管理

- GET /api/v1/users - 获取所有用户
- GET /api/v1/users/:id - 获取单个用户
- POST /api/v1/users - 创建用户
- PUT /api/v1/users/:id - 更新用户
- DELETE /api/v1/users/:id - 删除用户

## 运行项目

1. 安装依赖：
```bash
go mod download
```

2. 运行服务器：
```bash
go run main.go
```

3. 访问 API 文档：
```
http://localhost:8080/swagger/index.html
```

## 开发说明

- 使用 `swag init` 更新 Swagger 文档
- 遵循 RESTful API 设计规范
- 使用 Gin 框架处理 HTTP 请求 
