package grpc

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func NewServer(ctx context.Context) *Server {
	return &Server{
		ctx:      ctx,
		listener: nil,
		Server:   grpc.NewServer(),
	}
}

type Server struct {
	ctx      context.Context
	listener net.Listener
	Server   *grpc.Server
}

func (this *Server) SetListener(listener net.Listener) {
	this.listener = listener
}
func (this *Server) Run() {
	if this.listener == nil {
		panic("miss listener")
	}
	defer this.listener.Close()

	closeServer := make(chan bool)
	defer close(closeServer)
	go func() {
		c := make(chan os.Signal, 1)
		closeSignal := []os.Signal{
			syscall.SIGINT,  //Ctrl+C
			syscall.SIGTERM, //结束程序
			syscall.SIGQUIT, //Ctrl+/
		}
		signal.Notify(c, closeSignal...)
		<-c

		this.Server.GracefulStop()
		log.Println("grpcServer stop")
		closeServer <- true
	}()

	err := this.Server.Serve(this.listener)
	if err != nil {
		log.Fatalf("grpcServer close fail, err:%s", err.Error())
	} else {
		log.Println("grpcServer close")
	}
	<-closeServer
}
