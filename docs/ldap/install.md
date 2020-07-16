# ldap 安装


# Docker 安装

安装docker-compose尽量采用pip安装
```
yum install python-pip
pip install --upgrade pip
docker-compose up -d
```

采用Docker安装开发环境请参考: [docker-openldap](https://github.com/osixia/docker-openldap)


如果想简单使用可以:
```
docker run -p 389:389 -p 636:636 --name my-openldap-container --detach osixia/openldap:1.4.0
```

```
docker exec my-openldap-container ldapsearch -x -H ldap://localhost -b dc=example,dc=org -D "cn=admin,dc=example,dc=org" -w admin
```