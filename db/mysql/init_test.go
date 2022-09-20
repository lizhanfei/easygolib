package mysql

import (
	"fmt"
	"testing"
	"time"
)

func TestInitMysql(t *testing.T) {
	mysqlClient, err := InitMysql(&MysqlConf{
		DataBase:        "hyperf",
		Addr:            "127.0.0.1:3306",
		User:            "root",
		Password:        "123456",
		Charset:         "utf8mb4",
		MaxIdleConns:    5,
		MaxOpenConns:    20,
		ConnMaxIdlTime:  30 * time.Second,
		ConnMaxLifeTime: 600 * time.Second,
		ConnTimeOut:     1500 * time.Millisecond,
		WriteTimeOut:    1500 * time.Millisecond,
		ReadTimeOut:     1500 * time.Millisecond,
	}, nil)

	if err != nil {
		panic(fmt.Sprintf("init fail, err:%s", err))
	}

	mysqlClient.Table("brand_all").Create(map[string]interface{}{
		"brand_name": "aa",
	})
}
