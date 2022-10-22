package log

import (
	"errors"
)

type ConfFileWriter struct {
	Up              bool   `yaml:"Up"`
	LogPath         string `yaml:"LogPath"`
	RotateLogPath   string `yaml:"RotateLogPath"`
	WfLogPath       string `yaml:"WfLogPath"`
	RotateWfLogPath string `yaml:"RotateWfLogPath"`
}

type ConfConsoleWriter struct {
	Up    bool `yaml:"Up"`
	Color bool `yaml:"Color"`
}

type LogConfig struct {
	Level string            `yaml:"LogLevel"`
	FW    ConfFileWriter    `yaml:"FileWriter"`
	CW    ConfConsoleWriter `yaml:"ConsoleWriter"`
}

func SetupLogInstanceWithConf(lc LogConfig, logger *Logger) (err error) {
	if lc.FW.Up {
		if len(lc.FW.LogPath) > 0 {
			w := NewFileWriter()
			w.SetFileName(lc.FW.LogPath)
			w.SetPathPattern(lc.FW.RotateLogPath)
			w.SetLogLevelFloor(TRACE)
			if len(lc.FW.WfLogPath) > 0 {
				w.SetLogLevelCeil(INFO)
			} else {
				w.SetLogLevelCeil(ERROR)
			}
			logger.Register(w)
		}

		if len(lc.FW.WfLogPath) > 0 {
			wfw := NewFileWriter()
			wfw.SetFileName(lc.FW.WfLogPath)
			wfw.SetPathPattern(lc.FW.RotateWfLogPath)
			wfw.SetLogLevelFloor(WARNING)
			wfw.SetLogLevelCeil(ERROR)
			logger.Register(wfw)
		}
	}

	if lc.CW.Up {
		w := NewConsoleWriter()
		w.SetColor(lc.CW.Color)
		logger.Register(w)
	}
	switch lc.Level {
	case "trace":
		logger.SetLevel(TRACE)

	case "debug":
		logger.SetLevel(DEBUG)

	case "info":
		logger.SetLevel(INFO)

	case "warning":
		logger.SetLevel(WARNING)

	case "error":
		logger.SetLevel(ERROR)

	case "fatal":
		logger.SetLevel(FATAL)

	default:
		err = errors.New("Invalid log level")
	}
	return
}

func SetupDefaultLogWithConf(lc LogConfig) (err error) {
	defaultLoggerInit()
	return SetupLogInstanceWithConf(lc, logger_default)
}
