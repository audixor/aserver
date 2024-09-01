//
// Copyright (c) 2021-2024 Tenebris Technologies Inc.
// See LICENSE for further information.
//

package easysrv

import (
	"fmt"
	"log"
	"os"
)

// Logger is an interface that defines the logging methods and is compatible with log.Logger

type Logger interface {
	Debug(msg string, fields map[string]interface{})
	Info(msg string, fields map[string]interface{})
	Warn(msg string, fields map[string]interface{})
	Error(msg string, fields map[string]interface{})
	Fatal(msg string, fields map[string]interface{})
}

type Fields map[string]interface{}

func (e *EasySrv) SetLogger(logger Logger) {
	e.Logger = logger
}

type SimpleLogger struct {
	logger *log.Logger // A standard Go logger instance
}

// NewSimpleLogger creates a new SimpleLogger instance
func NewSimpleLogger(logFile string) (*SimpleLogger, error) {
	var logger *log.Logger

	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return &SimpleLogger{}, fmt.Errorf("error opening log file %s: %v", logFile, err)
		}
		logger = log.New(file, "", log.LstdFlags)
	} else {
		logger = log.New(os.Stdout, "", log.LstdFlags)
	}
	return &SimpleLogger{logger: logger}, nil
}

func (s SimpleLogger) Debug(msg string, fields map[string]interface{}) {
	s.WriteLog("DEBUG", msg, fields)
}

func (s SimpleLogger) Info(msg string, fields map[string]interface{}) {
	s.WriteLog("INFO", msg, fields)
}

func (s SimpleLogger) Warn(msg string, fields map[string]interface{}) {
	s.WriteLog("WARN", msg, fields)
}

func (s SimpleLogger) Error(msg string, fields map[string]interface{}) {
	s.WriteLog("ERROR", msg, fields)
}

func (s SimpleLogger) Fatal(msg string, fields map[string]interface{}) {
	s.WriteLog("FATAL", msg, fields)
}

// WriteLog reformats the event and sends to the standard Go logger
func (s SimpleLogger) WriteLog(level string, msg string, fields map[string]interface{}) {
	var txt = ""
	if fields != nil {
		for k, v := range fields {
			txt = txt + fmt.Sprintf(" %s=%v", k, v)
		}
	}
	s.logger.Printf("%s %s%s", level, msg, txt)
}
