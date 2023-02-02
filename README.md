# Golang System Service Scaffolding

1. 跨平台服务脚手架，需要Linux平台直接编译即可；
2. 终端启动打开默认浏览器；
3. 使用 logrus 格式化日志，rotatelogs 切割日志

``` sehll
go build -o .\bin\sys-service-scaffolding.exe .\main.go
```

```
2.1 设备初始化
    INIT_VIUINIT_REQ_SENDER_yyyyMMddHHmmsszzz.json
    SENDER：gantryId（19 位字符）+ cameraNum（3 位字符）
    示例：INIT_VIUINIT_REQ_G000511001000110100101_20190807100102345.json

2.2 基础数据上传
    SENDER：gantryId（19 位字符）+ cameraNum（3 位字符）
    示例：MON_BVIUBASEINFO_REQ_G000511001000110100101_20190807100102345.json

2.3 图片流水上传
    TRC_BVIU_REQ_G000511001000110100101_20190807100102345.json

2.4 图片上传
    TRC_BVIPU_REQ_G000511001000110100101_20190807100102345.json

2.5 状态信息上传
    MON_BVIUSTATE_REQ_G000511001000110100101_20190807100102345.json

```