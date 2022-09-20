package log

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func NewZLog(getRequestId func(ctx context.Context) string,
	getUriPath func(ctx context.Context) string,
	config LoggerConfig,
	encoder zapcore.Encoder) Zlog {

	res := &ZlogImpl{
		getRequestId: getRequestId,
		getUriPath:   getUriPath,
		module:       config.Module,
	}
	var stdLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= config.ZapLevel && lvl >= zapcore.DebugLevel
	})
	var normalLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= config.ZapLevel && lvl <= zapcore.InfoLevel
	})
	var errLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= config.ZapLevel && lvl >= zapcore.WarnLevel
	})
	if encoder == nil {
		encoderConfig := zap.NewProductionEncoderConfig() //指定时间格式
		encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.999999")
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
		encoder = zapcore.NewJSONEncoder(encoderConfig) //获取编码器,NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	}
	errWriteSyncer, _ := os.Create(config.Path + config.LogName + ".wf.log")        //日志文件存放目录
	normalWriteSyncer, _ := os.Create(config.Path + config.LogName + ".server.log") //日志文件存放目录

	var zapCore []zapcore.Core
	if config.Stdout {
		zapCore = append(zapCore, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), stdLevel))
	}

	if config.Log2File {
		//第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
		zapCore = append(zapCore, zapcore.NewCore(encoder, errWriteSyncer, errLevel))
		zapCore = append(zapCore, zapcore.NewCore(encoder, normalWriteSyncer, normalLevel))
	}
	logger := zap.New(zapcore.NewTee(zapCore...), zap.WithCaller(true)).WithOptions(zap.AddCallerSkip(1))

	res.zapLogger = logger
	res.sugaredLogger = res.zapLogger.Sugar()
	return res
}
