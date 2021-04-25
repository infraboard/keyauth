# 消息总线安装

keyauth支持2总消息总线: nats和kafka, 开发环境推荐使用docker部署:

+ nats 部署
```
docker run -p 4222:4222 -ti nats:latest
```