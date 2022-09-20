package log

import (
	"go.uber.org/zap/zapcore"
)

type LoggerConfig struct {
	ZapLevel zapcore.Level
	// 以下变量仅对开发环境生效
	Stdout   bool
	Log2File bool
	Path     string
	LogName  string
	Module   string
}
