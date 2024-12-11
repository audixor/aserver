//
// Copyright (c) 2024 Tenebris Technologies Inc.
// Please see the LICENSE file for details
//

package SimpleLogger

import (
	"fmt"
	"log"
	"os"
)

type SimpleLogger struct {
	logger *log.Logger // A standard Go logger instance
}

// Fields is a map of key-value pairs for additional log information
type Fields map[string]interface{}

// New creates a new SimpleLogger instance
func New(logFile string) (*SimpleLogger, error) {
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

func (s SimpleLogger) Debug(eid uint32, msg string, fields map[string]interface{}) {
	s.WriteLog("DEBUG", eid, msg, fields)
}

func (s SimpleLogger) Info(eid uint32, msg string, fields map[string]interface{}) {
	s.WriteLog("INFO", eid, msg, fields)
}

func (s SimpleLogger) Warning(eid uint32, msg string, fields map[string]interface{}) {
	s.WriteLog("WARNING", eid, msg, fields)
}

func (s SimpleLogger) Error(eid uint32, msg string, fields map[string]interface{}) {
	s.WriteLog("ERROR", eid, msg, fields)
}

func (s SimpleLogger) Fatal(eid uint32, msg string, fields map[string]interface{}) {
	s.WriteLog("FATAL", eid, msg, fields)
}

// WriteLog reformats the event and sends to the standard Go logger
func (s SimpleLogger) WriteLog(level string, eid uint32, msg string, fields map[string]interface{}) {
	var txt = ""
	if fields != nil {
		for k, v := range fields {
			txt = txt + fmt.Sprintf(" %s=%v", k, v)
		}
	}
	s.logger.Printf("%s %04d %s%s", level, eid, msg, txt)
}
