package log

import (
	"context"
	"go.uber.org/zap"
)

type Zlog interface {
	Debugf(ctx context.Context, format string, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
	Warnf(ctx context.Context, format string, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
	Fatalf(ctx context.Context, format string, args ...interface{})
}

type ZlogImpl struct {
	sugaredLogger *zap.SugaredLogger
	zapLogger     *zap.Logger
	getRequestId  func(ctx context.Context) string
	getUriPath    func(ctx context.Context) string
	module        string
}

func (this *ZlogImpl) getSugaredLogger(ctx context.Context) *zap.SugaredLogger {
	loggerRes := this.sugaredLogger.With(
		zap.String("requestId", this.getRequestId(ctx)),
		zap.String("appName", this.module),
		zap.String("uri", this.getUriPath(ctx),
		),
	)
	return loggerRes
}

func (this *ZlogImpl) Debugf(ctx context.Context, format string, args ...interface{}) {
	this.getSugaredLogger(ctx).Debugf(format, args...)
}

func (this *ZlogImpl) Infof(ctx context.Context, format string, args ...interface{}) {
	this.getSugaredLogger(ctx).Infof(format, args...)
}

func (this *ZlogImpl) Warnf(ctx context.Context, format string, args ...interface{}) {
	this.getSugaredLogger(ctx).Warnf(format, args...)
}

func (this *ZlogImpl) Errorf(ctx context.Context, format string, args ...interface{}) {
	this.getSugaredLogger(ctx).Errorf(format, args...)
}

func (this *ZlogImpl) Fatalf(ctx context.Context, format string, args ...interface{}) {
	this.getSugaredLogger(ctx).Fatalf(format, args...)
}
