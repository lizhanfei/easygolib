package redis

import "time"

type RedisConf struct {
	Addr            string
	Password        string
	DB              int
	MaxIdle         int
	MaxActive       int
	IdleTimeout     time.Duration
	MaxConnLifetime time.Duration
	ConnTimeOut     time.Duration
	ReadTimeOut     time.Duration
	WriteTimeOut    time.Duration
	BorrowSafeTime  time.Duration
}
