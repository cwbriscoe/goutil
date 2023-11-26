// Copyright 2023 Christopher Briscoe.  All rights reserved.
package logging

import (
	"io"
	"os"
	"path"
	"time"

	"github.com/natefinch/lumberjack/v3"
	"github.com/rs/zerolog"
)

// Logger contains a zerolog.Logger for logging functions and a
// lumberjack.Roller reference so that we can call Rotate() on
// our own schedule if desired
type Logger struct {
	*zerolog.Logger
	roller *lumberjack.Roller
}

// Nop is a no-operation logger instance with no roller attached
var (
	nop = zerolog.Nop()
	Nop = &Logger{
		Logger: &nop,
		roller: nil,
	}
)

// Config stores the information required to start a new logger
type Config struct {
	BaseDir    string        `json:"basedir"`
	FileName   string        `json:"filename"`
	MaxAge     time.Duration `json:"maxage"`
	MaxSize    int64         `json:"maxsize"`
	MaxBackups int           `json:"maxbackups"`
	Console    bool          `json:"console"`
	Compress   bool          `json:"compress"`
}

// NewLogger returns a new rolling logger based on the config parameters used
func NewLogger(cfg *Config) (*Logger, error) {
	var writers []io.Writer

	if cfg.Console {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}

	filename := path.Join(cfg.BaseDir, cfg.FileName)

	roller, err := lumberjack.NewRoller(filename, cfg.MaxSize, &lumberjack.Options{
		MaxAge:     cfg.MaxAge,
		MaxBackups: cfg.MaxBackups,
		Compress:   cfg.Compress,
	})
	if err != nil {
		return nil, err
	}

	writers = append(writers, roller)
	mw := io.MultiWriter(writers...)

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := zerolog.New(mw).With().Timestamp().Logger()

	return &Logger{
		Logger: &logger,
		roller: roller,
	}, nil
}

// Rotate will immediately close and rotate the current log file
func (l *Logger) Rotate() error {
	return l.roller.Rotate()
}
