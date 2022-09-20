package redis

import (
	"context"
	redigo "github.com/gomodule/redigo/redis"
	"time"
)

func InitClient(ctx *context.Context, conf *RedisConf, before func(ctx *context.Context, command string, args ...interface{}),
	after func(ctx *context.Context, startTime, endTime time.Time, err error, command string, args ...interface{})) *Client {
	p := &redigo.Pool{
		MaxIdle:         conf.MaxIdle,
		MaxActive:       conf.MaxActive,
		IdleTimeout:     conf.IdleTimeout,
		MaxConnLifetime: conf.MaxConnLifetime,
		Wait:            true,
		Dial: func() (conn redigo.Conn, e error) {
			con, err := redigo.Dial(
				"tcp",
				conf.Addr,
				redigo.DialPassword(conf.Password),
				redigo.DialConnectTimeout(conf.ConnTimeOut),
				redigo.DialReadTimeout(conf.ReadTimeOut),
				redigo.DialWriteTimeout(conf.WriteTimeOut),
			)
			if err != nil {
				return nil, err
			}
			con.Do("SELECT", conf.DB)
			return con, nil
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			timeSafe := time.Minute
			if conf.BorrowSafeTime > 0 {
				timeSafe = conf.BorrowSafeTime
			}
			if time.Since(t) < timeSafe {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}

	return &Client{
		ctx:    ctx,
		pool:   p,
		before: before,
		after:  after,
	}
}
