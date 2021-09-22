# go-mircoservice-learn
learn go mircoservice
- Deadlines

Deadlines 意指截止时间，在 gRPC 中强调 TL;DR（Too long, Don't read）并建议始终设定截止日期，为什么呢？
为什么要设置

当未设置 Deadlines 时，将采用默认的 DEADLINE_EXCEEDED（这个时间非常大）
如果产生了阻塞等待，就会造成大量正在进行的请求都会被保留，并且所有请求都有可能达到最大超时
这会使服务面临资源耗尽的风险，例如内存，这会增加服务的延迟，或者在最坏的情况下可能导致整个进程崩溃