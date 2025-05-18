# Personal Blog 后端

基于 Go + Gin + GORM + MySQL 的个人博客系统后端。

## 依赖安装

```bash
go mod tidy
```

## 数据库配置

默认使用 MySQL，连接信息在 `config/config.go` 中：

```
dsn := "root:password@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
```
请根据实际情况修改数据库用户名、密码和库名。

## 启动方式

```bash
go run main.go
```

## 目录结构

- main.go / 入口
- config/ 数据库配置
- middleware/ 中间件
- models/ GORM 模型
- routes/ 路由
- testCase/ 测试用例
- utils/ 共同工具