## 日志
基于uber-go/zap实现的简单日志操作类。

### 用法
```go
config := loggerConfig{
    Stdout:   true,//是否输出到控制台
    ZapLevel: zapcore.InfoLevel, //最低日志等级
    Path:     "./",//日志文件所在目录
    LogName:  "ttServer",//服务名称，会作为日志文件名的一部分
    Log2File: true,//是否输出到文件
}
logger := NewZLog(getRequestId, config, nil)
logger.Debugf(context.Background(), "debug name:%s", "aa")
logger.Infof(context.Background(), "info name:%s", "aa")
logger.Warnf(context.Background(), "warn name:%s", "aa")
logger.Errorf(context.Background(), "warn name:%s", "aa")
logger.Fatalf(context.Background(), "warn name:%s", "aa")
```
