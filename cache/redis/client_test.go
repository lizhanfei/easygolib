package redis

import (
	"context"
	"fmt"
	"github.com/alicebob/miniredis/v2"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestClient_Do(t *testing.T) {
	s := miniredis.RunT(t)
	assert := assert2.New(t)
	conf := &RedisConf{
		Addr:            s.Addr(),
		Password:        "",
		DB:              2,
		MaxIdle:         10,
		MaxActive:       20,
		IdleTimeout:     3 * time.Second,
		MaxConnLifetime: 6 * time.Second,
		ConnTimeOut:     1200 * time.Millisecond,
		ReadTimeOut:     1200 * time.Millisecond,
		WriteTimeOut:    1200 * time.Millisecond,
		BorrowSafeTime:  2 * time.Second,
	}

	ctx := context.Background()
	before := func(ctx *context.Context, command string, args ...interface{}) {
		return
	}
	after := func(ctx *context.Context, startTime, endTime time.Time, err error, command string, args ...interface{}) {
		if err != nil {
			fmt.Println(fmt.Sprintf("fail err:%s", err))
		} else {
			fmt.Println(fmt.Sprintf("cost time: %d", endTime.Sub(startTime)))
		}
	}
	client := InitClient(&ctx, conf, before, after)

	s.Select(conf.DB)
	client.Do("SET", "xo", "aa")
	val, err := s.Get("xo")
	assert.Nil(err)
	assert.Equal(val, "aa")
	time.Sleep(10 * time.Second)

	client.Do("SET", "xo", "bb")
	val, _ = s.Get("xo")
	assert.Equal(val, "bb")

	valOri, err := client.Do("GET", "xo")
	assert.Nil(err)
	assert.Equal(string(valOri.([]byte)), "bb")

	s.Close()
	_, err = client.Do("SET", "xo", "oo")
	assert.NotNil(err)

	client = InitClient(&ctx, conf, before, after)
	_, err = client.Do("SET", "xo", "oo")
	assert.NotNil(err)
}
