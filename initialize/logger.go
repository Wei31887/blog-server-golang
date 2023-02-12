package initialize

import (
	"blog/server/global"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type _zap struct{}
var Zap = new(_zap)

// Logger : initialize the log and return zap logger
func Logger() (logger *zap.Logger) {
	core := Zap.getLogCore()
	logger = zap.New(core)
	return
}

func filePathName() string {
	return filepath.Join(global.GLOBAL_CONFIG.Mylog.Path + "/" + global.GLOBAL_CONFIG.Mylog.Name + ".log")
}

// indicate the location of log file 
func (z *_zap) getLogWriter() zapcore.WriteSyncer {
	// Outputting log to console or file is depend on the config setting
	if global.GLOBAL_CONFIG.Mylog.Model == "console" {
		return zapcore.AddSync(os.Stdout)
	} else {
		file, _ := os.Create(filePathName())
		return zapcore.AddSync(file)
	}
}

// 
func (z *_zap) getEncoder() zapcore.Encoder {
	if global.GLOBAL_CONFIG.Mylog.Format == "JSON" {
		return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
}

// 
func (z *_zap) getCoreLevel() zapcore.Level {
	switch global.GLOBAL_CONFIG.Mylog.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

// 
func (z *_zap) getLogCore() zapcore.Core {
	return zapcore.NewCore(z.getEncoder(), z.getLogWriter(), z.getCoreLevel())
}


