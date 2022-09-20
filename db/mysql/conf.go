package mysql

import "time"

type MysqlConf struct {
	DataBase        string
	Addr            string
	User            string
	Password        string
	Charset         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxIdlTime  time.Duration
	ConnMaxLifeTime time.Duration
	ConnTimeOut     time.Duration
	WriteTimeOut    time.Duration
	ReadTimeOut     time.Duration
}
