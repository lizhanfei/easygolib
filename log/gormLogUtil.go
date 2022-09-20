package log

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

//用于gorm的log适配工具

type DbLog struct {
	LogLevel    logger.LogLevel
	logger      Zlog
	ckRequestID string
	appName     string
}

// LogMode 初始化DbLog
func (l *DbLog) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	l.ckRequestID = "requestId"
	return &newlogger
}
//SetLogger 设置logger 引擎
func (l *DbLog) SetLogger(zlog Zlog) {
	l.logger = zlog
}
//SetCkRequestId 设置获取requestId的协程key
func (l *DbLog) SetCkRequestId(newKey string) {
	l.ckRequestID = newKey
}
//SetAppName 设置日志中的appName 变量
func (l *DbLog) SetAppName(appName string) {
	l.appName = appName
}

// Info 日志
func (l DbLog) Info(ctx context.Context, msg string, data ...interface{}) {
	if nil != l.logger {
		l.logger.Infof(ctx, msg, data)
	}
}

//Warn 异常日志
func (l DbLog) Warn(ctx context.Context, msg string, data ...interface{}) {
	if nil != l.logger {
		l.logger.Warnf(ctx, msg, data)
	}
}

//Error 异常日志
func (l DbLog) Error(ctx context.Context, msg string, data ...interface{}) {
	if nil != l.logger {
		l.logger.Errorf(ctx, msg, data)
	}
}

//Trace trace 日志
func (l DbLog) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	end := time.Now()
	elapsed := end.Sub(begin)
	cost := float64(elapsed.Nanoseconds()/1e4) / 100.0

	// 请求是否成功
	msg := "mysql do success"
	resCode := 0
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// 没有找到记录不统计在请求错误中
		msg = err.Error()
		resCode = -1
	}

	sql, rows := fc()

	fields := l.commonFields(ctx)
	fields = append(fields,
		zap.String("sql", sql),
		zap.Int64("affectedrow", rows),
		zap.String("requestEndTime", l.GetFormatRequestTime(end)),
		zap.String("requestStartTime", l.GetFormatRequestTime(begin)),
		zap.Float64("cost", cost),
		zap.Int("resCode", resCode),
	)

	l.logger.Infof(ctx, msg, fields)
}

//commonFields 日志中的常规变量（requestId、prot、module）
func (l DbLog) commonFields(ctx context.Context) []zap.Field {
	var requestID string
	if c, ok := ctx.(*gin.Context); ok && c != nil {
		requestID, _ = c.Value(l.ckRequestID).(string)
	}
	fields := []zap.Field{
		zap.String("requestId", requestID),
		zap.String("prot", "mysql"),
		zap.String("module", l.appName),
	}
	return fields
}

//GetFormatRequestTime 计算sql执行时间
func (l DbLog) GetFormatRequestTime(time time.Time) string {
	return fmt.Sprintf("%d.%06d", time.Unix(), time.Nanosecond()/1e3)
}