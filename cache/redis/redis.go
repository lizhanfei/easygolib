package redis

import (
	"context"
	redigo "github.com/gomodule/redigo/redis"
	"time"
)

type Client struct {
	pool   *redigo.Pool
	ctx    *context.Context
	before func(ctx *context.Context, command string, args ...interface{})
	after  func(ctx *context.Context, startTime, endTime time.Time, err error, command string, args ...interface{})
}

func (this *Client) Do(command string, args ...interface{}) (res interface{}, err error) {
	startTime := time.Now()

	this.before(this.ctx, command, args...)
	c := this.pool.Get()
	if c.Err() != nil {
		this.after(this.ctx, startTime, time.Now(), c.Err(), command, args...)
		return nil, c.Err()
	}

	replay, err := c.Do(command, args...)
	if errClose := c.Close(); errClose != nil {
		this.after(this.ctx, startTime, time.Now(), errClose, command, args...)
		return nil, errClose
	}
	if err != nil {
		this.after(this.ctx, startTime, time.Now(), err, command, args...)
		return nil, err
	}
	this.after(this.ctx, startTime, time.Now(), nil, command, args...)

	return replay, nil
}
