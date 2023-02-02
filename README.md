# Golang System Service Scaffolding

1. 跨平台服务脚手架，需要Linux平台直接编译即可；
2. 终端启动打开默认浏览器；
3. 使用 logrus 格式化日志，rotatelogs 切割日志

``` sehll
go build -o .\bin\sys-service-scaffolding.exe .\main.go
```