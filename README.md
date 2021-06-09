## QuickStart

### 依赖
mysql 
redis
nsq
aliyun sls
qq mail

### config

```yml
# 配置示例
config.example.yml
```

### Docker

```bash

# 创建配置文件并启动应用
cd /demo
touch config.yml

docker run --name dora -d -p 8221:8222 \
      -v /demo/config.yml:/internal/config.yml \ 
      -it nancode/dora:latest

docker run --name dora -d -p 8221:8222 \
    -v /root/dora/tmp:/internal/tmp \
    -v /root/dora/config.yml:/internal/config.yml \
    -it nancode/dora:latest

docker run --name dora -d  --network host  -v /root/dora/tmp:/internal/tmp \
     -v /root/dora/config.yml:/internal/config.yml \
     -it registry.cn-hangzhou.aliyuncs.com/nancode/dora:latest

# 更新正在运行的 container
docker run --rm \
    -v /var/run/docker.sock:/var/run/docker.sock \
    containrrr/watchtower -cR \
    dora --debug
```

### plan

sdk -> mq -> 阿里云sls -> mysql

```
用户模块
监控模块
打点模块

# 非重点模块延后
```

## Contributing

## Author

- [nan](https://github.com/nanzm)

## License