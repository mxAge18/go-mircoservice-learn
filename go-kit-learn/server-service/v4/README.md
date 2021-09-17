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
  

## v4 重构 command line 启动
### 利用第三方包进行命令行处理
- github.com/spf13/cobra
- 手动输入服务名和端口号，启动微服务并注册到consul

### 限流
- 请求数据量限制，返回自定义错误429
- rate包：
  - rate.NewLimiter(r, b)
  - Wait/WaitN
  - Allow/AllowN
- 通过加endpoint 中间件的方式集成到请求中。达到限流目的


### error自定义
- 自定义error struct实现Error方法，达到自定义返回http状态🐴+ 错误message的目的
### 熔断 降级
https://github.com/afex/hystrix-go
```go
hystrix.ConfigureCommand("my_command", hystrix.CommandConfig{
Timeout:               1000,
MaxConcurrentRequests: 100,
ErrorPercentThreshold: 25,
})

hystrix.Go("my_command", func() error {
// talk to other services
return nil
}, func(err error) error {
// do this when services are down
return nil
})

```
#### 服务熔断
- 如果调用微服务超时，进行服务熔断处理。超时返回错误，不等待微服务的响应
- 最大并发限制
#### 服务降级
发生服务熔断时，调用客户端本地缓存的数据。

