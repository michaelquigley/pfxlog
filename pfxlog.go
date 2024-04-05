package pfxlog

import (
	"context"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"log/slog"
	"os"
	"runtime"
	"time"
)

func GlobalInit(level slog.Level, options *Options) {
	if defaultEnv("PFXLOG_NO_JSON", false) || terminal.IsTerminal(int(os.Stdout.Fd())) {
		logger := slog.New(NewPrettyHandler(level, options))
		slog.SetDefault(logger)
	} else {
		logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level, AddSource: true}))
		slog.SetDefault(logger)
	}
	globalOptions = options
}

func Logger() Builder {
	return Builder{slog.Default()}
}

func ContextLogger(context string) Builder {
	return Builder{slog.Default().With(slog.String("_context", context))}
}

func Info(args ...interface{}) {
	if !slog.Default().Enabled(context.Background(), slog.LevelInfo) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelInfo, fmt.Sprint(args...), pcs[0])
	_ = slog.Default().Handler().Handle(context.Background(), r)
}

func Infof(format string, args ...interface{}) {
	if !slog.Default().Enabled(context.Background(), slog.LevelInfo) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelInfo, fmt.Sprintf(format, args...), pcs[0])
	_ = slog.Default().Handler().Handle(context.Background(), r)
}

type Builder struct {
	*slog.Logger
}

func (b Builder) Infof(format string, args ...interface{}) {
	if !b.Enabled(context.Background(), slog.LevelInfo) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelInfo, fmt.Sprintf(format, args...), pcs[0])
	_ = b.Handler().Handle(context.Background(), r)
}

var globalOptions *Options
