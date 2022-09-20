## httpServer
基于gin的httpServer，支持优雅关机。

### 优雅关机
监听 syscall.SIGINT、syscall.SIGTERM、syscall.SIGQUIT 信号关机；
关机时，等待未完成的请求完成，可以指定最长等待时间

### 用法
```go
tcpAddress := ":18880"
l, _ := NewListenerTcp(tcpAddress)
readTimeout := 15 * time.Second //httpserver 读超时时间
writeTimeout := 15 * time.Second //httpserver写超时时间
closeWaitTimeout := 0  //关机最长等待时间，传0默认为30s
s := NewServer(ginEngin, l, readTimeout, writeTimeout, closeWaitTimeout)
s.Run()
```