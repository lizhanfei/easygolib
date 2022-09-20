package log

import (
	"context"
	"go.uber.org/zap/zapcore"
	"strconv"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	config := LoggerConfig{
		Stdout:   true,
		ZapLevel: zapcore.InfoLevel, //最低日志等级
		Path:     "./",
		LogName:  "ttServer",
		Log2File: true,
	}
	logger := NewZLog(getRequestId, getUriPath, config, nil)
	logger.Debugf(context.Background(), "debug name:%s", "aa")
	logger.Infof(context.Background(), "info name:%s", "aa")
	logger.Warnf(context.Background(), "warn name:%s", "aa")
	logger.Errorf(context.Background(), "warn name:%s", "aa")
	//logger.FatalF(context.Background(), "warn name:%s", "aa")
}

func getRequestId(ctx context.Context) string {
	usec := uint64(time.Now().UnixNano())
	return strconv.FormatUint(usec&0x7FFFFFFF|0x80000000, 10)
}

func getUriPath(ctx context.Context) string {
	return "127.0.0.1"
}
