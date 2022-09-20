package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
	"time"
)

func TestServerBase(t *testing.T) {
	ginEngin := gin.New()
	l, _ := NewListenerTcp(":8080")

	s := NewServer(ginEngin, l, 10, 10, 0)
	s.Run()
}


func TestServerHuman(t *testing.T) {
	ginEngin := gin.New()
	ginEngin.GET("/test", func(context *gin.Context) {
		context.JSONP(http.StatusOK, map[string]string{
			"data": "success",
		})
	})
	l, _ := NewListenerTcp(":8080")

	s := NewServer(ginEngin, l, 10 * time.Millisecond, 10 * time.Millisecond, 0)
	s.Run()
}

func TestServerHumanTimeLong(t *testing.T) {
	ginEngin := gin.New()
	ginEngin.GET("/test", func(context *gin.Context) {
		time.Sleep(10*time.Second)
		context.JSONP(http.StatusOK, map[string]string{
			"data": "after sleep success",
		})
	})
	l, _ := NewListenerTcp(":8081")

	s := NewServer(ginEngin, l, 15 * time.Second, 15 * time.Second, 0)
	s.Run()
}