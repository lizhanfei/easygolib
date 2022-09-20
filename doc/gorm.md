## gorm
这里对gorm初始化做一些封装，另外对gorm的日志提供适配

### 使用
```go
mysqlClient, err := InitMysql(&MysqlConf{
    DataBase:        "hyperf",
    Addr:            "127.0.0.1:3306",
    User:            "root",
    Password:        "123456",
    Charset:         "utf8mb4",
    MaxIdleConns:    5,//连接池 最大空闲连接数
    MaxOpenConns:    20,//连接池最大连接数
    ConnMaxIdlTime:  30 * time.Second,//连接池连接最大空闲周期
    ConnMaxLifeTime: 600 * time.Second,//连接池中连接最大生命周期
    ConnTimeOut:     1500 * time.Millisecond,
    WriteTimeOut:    1500 * time.Millisecond,
    ReadTimeOut:     1500 * time.Millisecond,
}, &log.DbLog{})

//这里的 &log.DbLog{}) 实现在 ./log/gormLogUtil.go 中
```