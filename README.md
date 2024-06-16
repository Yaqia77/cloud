# Air
特点：

1. 监视代码变化并自动重启应用程序
2. 友好的配置文件支持
3. 快速安装和易于使用
   
## 安装

```
go install github.com/air-verse/air@latest
```

## 使用
运行 Air：
```
air
```

Air 默认监视当前目录下所有 go 文件的变化，并自动重启应用程序。
Air 会生成一个默认的配置文件 .air.toml，你可以根据需要进行修改