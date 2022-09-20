## grpcServer
基于 google.golang.org/grpc，增加了信号关机

### 优雅关机
监听 syscall.SIGINT、syscall.SIGTERM、syscall.SIGQUIT 信号关机；
grpc 本身会在关机时，向客户端推送断开命令

### 用法
```go
tcpAddress := ":19090"
ctxBase := context.Background()
grpcServer := NewServer(ctxBase)
//注册grpc服务
XX.RegisterUserServiceServer(grpcServer, XXX{})
lis, err := grpc.NewListenerTcp(tcpAddress)
if err != nil {
    panic("grpc listen fail")
}
grpcServer.SetListener(lis)
grpcServer.Run()
```