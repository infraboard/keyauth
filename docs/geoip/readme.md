# GEOIP

登录日志审计时需要解析IP对应的城市, 所以需要


## 什么是GepIP
所谓GeoIP，就是通过来访者的IP， 定位他的经纬度，国家/地区，省市，甚至街道等位置信息。这里面的技术不算难题，关键在于有个精准 的数据库。
有了准确的数据源就奇货可居赚点小钱，可是发扬合作精神，集体贡献众人享用是我们追求的。

 

## GeoIP如 何使用
首先我们需要数据信息，所以先获取一个免费的数据库：GeoIP.dat.gz ，接着解压得到：GeoIP.dat， 然后就是对数据文件的按需操作。

到[这个页面](https://dev.maxmind.com/geoip/geoip2/geolite2/)下载ip库

下载需要注册账号, 注册后，选择Download Files, 下载名为GeoLite2 City的文件

这里没有使用官方的SDK, 官方的做法是直接发文件载入内存中, 然后查询

我打算将数据导入MongoDB, 然后查询, 但是由于数据 不是直接提供的IP地址，而且提供的CIDR,
所以, 需要技术出cidr的开始和结束的IP, 并且报错为int类型方便比较大小(通过比较开始和结束从而判断ip属于哪个cidr)
```
"_id" : "31.11.43.0/24",
"ipType" : "CIDR",
"first" : "31.11.43.0",
"last" : "31.11.43.255",
"start" : NumberLong(520825600),
"end" : NumberLong(520825855),
"count" : NumberLong(256)
}
```

## 另一个选择 GeoIP-CN
```
https://r.coding-space.cn/r/4341 最小巧、最准确、最实用的中国大陆IP段+GeoIP2数据库
```

## 验证
可以到 [iP地址归属地查询](https://www.ip138.com/) 查询IP地址然后比对, 判断地址库正确性

## 参考

+ [使用geoip库, 通过ip获取国家，省市，城市](http://www.fecmall.com/topic/806)
+ [导入数据到MongoDB, 实现查询](http://www.kode12.com/kode12/mongodb/store-compare-ip-address-mongodb-study/)