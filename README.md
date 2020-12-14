# 5GC API 服务器

## 环境需求
- Go 1.15.4 或更新的版本
- openapi-generator 5.0.0-beta3 或更新的版本

Edgecore 5GC API 服务器. 更多关于 5GC API 的细节可以参考
[5GC API 文件](https://app.swaggerhub.com/apis-docs/breezestars/5gc/1.0.0).


## 概况
- API version: 1.0.0
- Build date: 2020-12-14T01:12:40.970586+08:00[Asia/Taipei]

### 执行服务器
要执行服务器，务必确认同层目录下有 config.json

config.json 内包含模拟用的资料

#### 执行
```
// 未编译的状况下用 go 直接跑
go run main.go

or

// 执行已编译的档案
./5gc-api-gw
```


