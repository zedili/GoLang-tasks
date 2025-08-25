# Blog API 项目

基于 Go 和 Gin 框架构建的简单博客系统 API，提供用户认证、文章管理和评论功能。

## 主要依赖库

- **Gin**: Web 框架 (github.com/gin-gonic/gin)
- **GORM**: ORM 库 (gorm.io/gorm)
- **MySQL 驱动**: GORM 的 MySQL 适配器 (gorm.io/driver/mysql)
- **JWT**: 认证令牌 (github.com/dgrijalva/jwt-go)
- **数据验证**: 请求参数验证 (github.com/go-playground/validator/v10)

## 核心功能

- ✅ 用户注册和登录 (JWT 认证)
- ✅ 文章创建和管理
- ✅ 评论系统
- ✅ RESTful API 设计
- ✅ MySQL 数据库存储
- ✅ 文件日志记录

## 快速开始

```bash
# 安装依赖
go mod download

# 启动服务
go run main.go
```

服务将在 `http://localhost:8090` 启动，需要提前配置 MySQL 数据库连接信息。