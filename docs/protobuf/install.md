# 安装protoc

1. 下载protoc 版本1.39.1, 注意版本一定要对上 [protoc下载地址](https://github.com/protocolbuffers/protobuf/releases)
```
# 1.安装protoc编译器,  项目使用版本: v3.19.1
# 下载预编译包安装: https://github.com/protocolbuffers/protobuf/releases
```

2. 将需要的文件手动copy到对于地方
```sh
cp protoc-3.19.1-osx-x86_64/bin/protoc /usr/local/bin
cp protoc-3.19.1-osx-x86_64/include/*  /usr/local/include
```

3. 安装gprc相关插件
```sh
# 1.protoc-gen-go go语言查询, 项目使用版本: v1.27.1   
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# 2.安装protoc-gen-go-grpc插件, 项目使用版本: 1.1.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

4. 安装自定义Tag插件
```
# 1.安装自定义proto tag插件
go install github.com/favadi/protoc-go-inject-tag@latest
```

5. 安装项目依赖的protobuf
```
cp -r docs/include/github.com /usr/local/include
```
