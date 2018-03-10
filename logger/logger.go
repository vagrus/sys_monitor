package logger

import (
	"log"
	"os"
	"errors"
	"strings"
)

const (
	_ uint			= iota
	LEVEL_DEBUG
	LEVEL_INFO
	LEVEL_WARN
	LEVEL_ERROR
	LEVEL_FATAL
)

var levelPres = map[uint]string{
	LEVEL_DEBUG: "DEBUG",
	LEVEL_INFO: "INFO",
	LEVEL_WARN: "WARNING",
	LEVEL_ERROR: "ERROR",
	LEVEL_FATAL: "FATAL",
}

var logDir = "/var/log/app"
var logFilename = "sys_monitor.log"
var logLevel uint

var logger *log.Logger
var logFile *os.File

func Init(level uint) error {
	var err error

	if level > LEVEL_FATAL {
		return errors.New("wrong log level")
	}
	logLevel = level

	fullFilename := logDir + "/" + logFilename
	logFile, err = os.OpenFile(fullFilename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	log.SetOutput(logFile)
	logger = log.New(logFile, "", log.LstdFlags | log.LUTC)

	return nil
}

func Stop() {
	if logFile != nil {
		logFile.Close()
	}
}

func Debug(vars ...interface{}) {
	if logLevel <= LEVEL_DEBUG {
		write(LEVEL_DEBUG, vars)
	}
}

func Warn(vars ...interface{}) {
	if logLevel <= LEVEL_WARN {
		write(LEVEL_WARN, vars)
	}
}

func Error(vars ...interface{}) {
	if logLevel <= LEVEL_ERROR {
		write(LEVEL_ERROR, vars)
	}
}

func write(level uint, vars []interface{}) {
	var output []string
	for _, v := range vars {
		strV, err := ForLog(v)
		if err != nil {
			output = append(output, "\"", err.Error(), "\"")
		} else {
			output = append(output, strV)
		}
	}

	logger.Printf("[%s] %s", levelPres[level], strings.Join(output, " "))
}
