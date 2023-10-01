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

// Config stores the information required to start a new logger
type Config struct {
	BaseDir    string
	FileName   string
	MaxAge     time.Duration
	MaxSize    int64
	MaxBackups int
	Console    bool
	Compress   bool
}

// NewLogger returns a new rolling logger based on the config parameters used
func NewLogger(config Config) (*Logger, error) {
	var writers []io.Writer

	if config.Console {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}

	filename := path.Join(config.BaseDir, config.FileName)

	roller, err := lumberjack.NewRoller(filename, config.MaxSize, &lumberjack.Options{
		MaxAge:     config.MaxAge,
		MaxBackups: config.MaxBackups,
		Compress:   config.Compress,
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
