# Analog-network
- 网络延迟与网络丢包模拟工具
- 多数据中心或集群环境下测试集群可靠性，通过控制L4层读写字节流实现模拟多数据中心或集群环境下网络延迟与网络丢包

# Quick start

## config.json 中配置多个远程目标IP端口与本地端口
``` 
      "port": "1000",               --本地访问端口
      "target": "127.0.0.1:3306",   --远程目标
      "desc": "mysql",
      "ctl": {
        "in": {
          "lossRate": 0,            --丢包率百分比，整数
          "delayMs": 0              --延迟(毫秒)
        },
        "out": {
          "lossRate": 0,
          "delayMs": 0
        }
      }
```
## 启动
 
`` 
./analog-network-linux
``
## 更改 config.json 中延迟与丢包值后生效

``` 
curl localhost:9999/apply
```


