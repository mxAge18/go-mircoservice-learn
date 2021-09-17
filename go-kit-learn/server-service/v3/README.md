# v1 简单的获取UserName的http服务

## demo目录
```bin/bash
service
--user-endpoint.go // 定义endpoint
--user-service.go // 定义service
--user-transport.go // 定义transport
main.go
```

## 基于go-kit构建
https://gokit.io/

https://github.com/go-kit/kit

## v2 RESTFul api /user/

## v3 开启服务的同时注册服务到consul
### 重构了相关代码

### 优雅的退出（退出时deregister 服务）
  优雅退出的方式是通过os.Signal的监听进行，通过channel通信，监听到管道中有退出或者错误信号，退出服务
  signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
  

## v4 command-line 注册一个服务（模拟一个服务多个端口）负载均衡
- flag包可以监听输入
- 之前学过另外的第三方包：github.com/spf13/cobra 使用这个包进行命令行的开发