package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//NewServer 实例化server
func NewServer(engine *gin.Engine, listener net.Listener, readTimeout time.Duration,
	writeTimeout time.Duration, closeWaitTimeout time.Duration) Server {
	if closeWaitTimeout <= 0 {
		closeWaitTimeout = 30 * time.Second
	}
	return Server{
		engine:           engine,
		listener:         listener,
		readTimeout:      readTimeout,
		writeTimeout:     writeTimeout,
		closeWaitTimeout: closeWaitTimeout,
	}
}

type Server struct {
	engine           *gin.Engine
	listener         net.Listener
	readTimeout      time.Duration
	writeTimeout     time.Duration
	closeWaitTimeout time.Duration
}

//Run 启动server服务
func (this *Server) Run() {
	if this.listener == nil {
		panic("miss listener")
	}
	defer this.listener.Close()

	httpServer := &http.Server{
		Handler: this.engine,
	}
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

		ctx, cancel := context.WithTimeout(context.Background(), this.closeWaitTimeout)
		defer cancel()
		if err := httpServer.Shutdown(ctx); err != nil {
			log.Fatalf("httpServer shutdown fail, err:%s", err.Error())
		} else {
			log.Println("httpServer shutdown")
		}
		closeServer <- true
	}()
	if this.readTimeout > 0 {
		httpServer.ReadHeaderTimeout = this.readTimeout
	}
	if this.writeTimeout > 0 {
		httpServer.WriteTimeout = this.writeTimeout
	}
	if err := httpServer.Serve(this.listener); err != http.ErrServerClosed && err != nil {
		log.Fatalf("httpServer close fail, err:%s", err.Error())
	} else {
		log.Println("httpServer close")
	}
	<-closeServer
}
