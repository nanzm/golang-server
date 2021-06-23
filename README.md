# dora-server
## 启动依赖
```shell
# 一键启动 mysql redis nsq  
make base-up

# 查看 logs
make base-logs

# 关闭
make base-down

# 清除数据
make data-clean
```

## 本地运行
> 配置示例 config.example.yml
```shell
# 创建配置文件
cp config.example.yml config.yml

# 填入你的 mail、 slsLog 或 elastic 配置
vi config.yml
```
```shell
# transit 服务
go run cmd/transit/main.go

# manage 服务
go run cmd/manage/main.go
```

