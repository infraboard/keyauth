# mongoDB 安装


# 如何安装

##  repo安装
参考官方按照文档 [官方按照文档](https://docs.mongodb.com/manual/installation/)

如果安装较慢 则采用淘宝源进行安装
编辑 /etc/yum.repos.d/mongodb-org-4.4.repo 添加如下内容
```
[mongodb-org-4.4]
name=MongoDB Repository
baseurl=https://mirrors.aliyun.com/mongodb/yum/redhat/$releasever/mongodb-org/4.4/x86_64/
gpgcheck=1
enabled=1
gpgkey=https://www.mongodb.org/static/pgp/server-4.4.asc
```

启动服务
```
systemctl start mongod
```

查看服务状态
```
systemctl status mongod
```

开机启动
```
systemctl enable mongod
```

## docker 安装

```
docker pull mongo
docker run -itd -p 27017:27017 --name mongo mongo
```

# 创建管理用户

进入Mongo Shell
```
docker exec -it mongo mongo
```

创建管理员账号
```
use admin
db.createUser({user:"admin",pwd:"123456",roles:["root"]})
db.auth("admin", "123456")
```

修改服务监听地址
```
# network interfaces
net:
  port: 17232
  bindIp: 0.0.0.0  # Enter 0.0.0.0,:: to bind to all IPv4 and IPv6 addresses or, alternatively, use the net.bindIpAll setting.
```


编辑 /etc/mongod.conf 开启认证访问
```
security:
  authorization: enabled
```

重启服务
```
systemctl restart mongod
```

如何修改密码
```
db.changeUserPassword("admin", "xxxx");
```

# 添加库用户
```
use keyauth
db.createUser({user: "keyauth", pwd: "123456", roles: [{ role: "dbOwner", db: "keyauth" }]})
```
