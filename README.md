# api-gw

## 环境需求
- Go 1.15.4 or later

这个项目目的为协助前端开发时，有一个 PISC API GW 模拟器可以测试使用

pisc-api-gw 启动时，可针对 port :8080 进行 API 操作

API 文件请参阅

https://app.swaggerhub.com/apis-docs/breezestars/pisc-proto/1.0

### Execute
```
// 执行已编译档
./pisc-api-gw
```

### Command
```
// 编译模拟器，产生只能于 Linux 上执行的 pisc-api-gw
make

// 编译 proto 产生需要的代码
make proto

// 清除 pisc-api-gw 与 proto 产生的代码
make clean
```