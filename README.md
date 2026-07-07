# Gin 框架 oneof 验证标签对比示例

本项目通过两个对比示例，直观展示 Gin 框架中 `oneof` 验证标签不同写法的差异。

## 项目结构

```
├── README.md
├── cmd
│   ├── gin-binding-grace      # 优雅的实现
│   │   └── main.go
│   └── gin-binding-oops       # 有问题的实现
│       └── main.go
├── go.mod
├── go.sum
└── internal
    └── pkg
        └── core
            └── response.go    # 通用响应结构
```

## 快速开始

### 1. 启动服务

```bash
# 启动有问题的实现（端口 8001）
go run cmd/gin-binding-oops/main.go

# 启动优雅的实现（端口 8002）
go run cmd/gin-binding-grace/main.go
```

### 2. 测试接口

```bash
# 测试有问题的实现
curl "http://127.0.0.1:8001/api/v1/notifications?type=whatsapp"

# 测试优雅的实现
curl "http://127.0.0.1:8002/api/v1/notifications?type=whatsapp"
```

## 对比说明

| 对比项 | gin-binding-oops | gin-binding-grace |
|--------|------------------|-------------------|
| 写法 | `oneof='email''sms''whatsapp''slack'''` | `oneof=all email sms whatsapp slack` |
| 可读性 | ❌ 极差 | ✅ 清晰 |
| 可维护性 | ❌ 难以维护 | ✅ 易于维护 |
| 空值处理 | ❌ 语义不明 | ✅ `default=all` |
| 新增类型 | ❌ 不知道如何添加 | ✅ 直接在末尾添加 |

## 技术栈

- Go 1.21+
- Gin v1.10+
- go-playground/validator v10

## 相关文章

详细分析请参考：[Gin 框架 oneof 验证标签：单引号写法的陷阱与最佳实践](https://fishfianl/posts/gin/oneof-single-quote-misunderstanding.html)

## License

MIT
