package initialize

import (
	"blog/server/global"
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type log struct{}
var Log = new(log)

// InitializeLogger : initialize the log and return zap logger
func InitializeLogger() (logger *zap.Logger) {
	// Outputting log to console or file is depend on the config setting
	core := Log.getLogCore()
	logger = zap.New(core)
	return
}

// indicate the location of log file 
func (l log) getLogWriter() zapcore.WriteSyncer {
	if global.GLOBAL_CONFIG.Mylog.Model == "console" {
		return zapcore.AddSync(os.Stdout)
	} else {
		file, _ := os.Create(filePathName())
		return zapcore.AddSync(file)
	}
}

func filePathName() string {
	fmt.Println(filepath.Join(global.GLOBAL_CONFIG.Mylog.Path + global.GLOBAL_CONFIG.Mylog.Name + ".log"))
	return filepath.Join(global.GLOBAL_CONFIG.Mylog.Path + "/" + global.GLOBAL_CONFIG.Mylog.Name + ".log")
}

// 
func (l log) getEncoder() zapcore.Encoder {
	if global.GLOBAL_CONFIG.Mylog.Format == "JSON" {
		return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
}

// 
func (l log) getCoreLevel() zapcore.Level {
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
func (l log) getLogCore() zapcore.Core {
	return zapcore.NewCore(l.getEncoder(), l.getLogWriter(), l.getCoreLevel())
}


