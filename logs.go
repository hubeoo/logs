package logs

import (
	"io"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *log.Logger

func InitLog(appName string) {
	// Create the full path for the new log file.
	logFileName := filepath.Join("/home/mydocker/log", appName+".log")

	// Set up logging with lumberjack for log file rotation.
	logFile := &lumberjack.Logger{
		Filename:   logFileName, // Log file path
		MaxSize:    10,          // Max size in megabytes before log file is rotated
		MaxBackups: 3,           // Maximum number of old log files to retain
		MaxAge:     60,          // Max age in days to retain old log files
		Compress:   false,       // Compress/zip old log files
		LocalTime:  true,
	}

	multiWriter := io.MultiWriter(logFile, os.Stderr)
	logger = log.New(multiWriter)

	logger.SetLevel(log.DebugLevel)

	logger.SetPrefix(appName)
	logger.SetReportTimestamp(true)
	logger.SetTimeFormat("2006-01-02 15:04:05")
	logger.SetCallerOffset(1)
	logger.SetReportCaller(true)
}

func Fatal(format string, values ...any) {
	if len(values) == 0 {
		logger.Fatal(format)
	} else {
		logger.Fatalf(format, values...)
	}
}
func Info(format string, values ...any) {
	if len(values) == 0 {
		logger.Info(format)
	} else {
		logger.Infof(format, values...)
	}
}

func Warn(format string, values ...any) {
	if len(values) == 0 {
		logger.Warn(format)
	} else {
		logger.Warnf(format, values...)
	}
}

func Debug(format string, values ...any) {
	if len(values) == 0 {
		logger.Debug(format)
	} else {
		logger.Debugf(format, values...)
	}
}
func Error(format string, values ...any) {
	if len(values) == 0 {
		logger.Error(format)
	} else {
		logger.Errorf(format, values...)
	}
}
