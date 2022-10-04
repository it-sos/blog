# vscode 开发运行代码
1. 安装coder runner插件后，直接运行 ctrl+alt+n，在本项目.vscode settings.json中可见配置项
1. 将backend作为工作区加入，可以解决无法正确调用go.mod的问题

# 迁移数据表
```sql
create database blog_dev charset=utf8mb4 collate=utf8mb4_general_ci;
```

# 同步数据表
go run migration/main.go

# 创建初始化用户和密码 admin / 123456
go test -timeout 30s -run ^Test_authService_Register$ github.com/it-sos/blog/services


# 安装 rizla
go install github.com/kataras/rizla
