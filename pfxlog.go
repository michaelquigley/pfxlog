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

const ChannelKey = "_channel"

func GlobalInit(level slog.Level, options *Options) {
	var handler slog.Handler
	if defaultEnv("PFXLOG_NO_JSON", false) || terminal.IsTerminal(int(os.Stdout.Fd())) {
		handler = NewPrettyHandler(level, options)
		if options.Handler != nil {
			handler = options.Handler
		}
		logger := slog.New(handler)
		slog.SetDefault(logger)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level, AddSource: true})
		if options.Handler != nil {
			handler = options.Handler
		}
		logger := slog.New(handler)
		slog.SetDefault(logger)
	}
	globalOptions = options
}

func Logger() Builder {
	return Builder{slog.Default()}
}

func Debug(args ...interface{}) {
	if !slog.Default().Enabled(context.Background(), slog.LevelDebug) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelDebug, fmt.Sprint(args...), pcs[0])
	_ = slog.Default().Handler().Handle(context.Background(), r)
}

func Debugf(format string, args ...interface{}) {
	if !slog.Default().Enabled(context.Background(), slog.LevelDebug) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelDebug, fmt.Sprintf(format, args...), pcs[0])
	_ = slog.Default().Handler().Handle(context.Background(), r)
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

func Warn(args ...interface{}) {
	if !slog.Default().Enabled(context.Background(), slog.LevelWarn) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelWarn, fmt.Sprint(args...), pcs[0])
	_ = slog.Default().Handler().Handle(context.Background(), r)
}

func Warnf(format string, args ...interface{}) {
	if !slog.Default().Enabled(context.Background(), slog.LevelWarn) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelWarn, fmt.Sprintf(format, args...), pcs[0])
	_ = slog.Default().Handler().Handle(context.Background(), r)
}

func Error(args ...interface{}) {
	if !slog.Default().Enabled(context.Background(), slog.LevelError) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelError, fmt.Sprint(args...), pcs[0])
	_ = slog.Default().Handler().Handle(context.Background(), r)
}

func Errorf(format string, args ...interface{}) {
	if !slog.Default().Enabled(context.Background(), slog.LevelError) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelError, fmt.Sprintf(format, args...), pcs[0])
	_ = slog.Default().Handler().Handle(context.Background(), r)
}

type Builder struct {
	*slog.Logger
}

func Channel(channel string) Builder {
	return Builder{slog.Default().With(slog.String(ChannelKey, channel))}
}

func (b Builder) Debugf(format string, args ...interface{}) {
	if !b.Enabled(context.Background(), slog.LevelDebug) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelDebug, fmt.Sprintf(format, args...), pcs[0])
	_ = b.Handler().Handle(context.Background(), r)
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

func (b Builder) Warnf(format string, args ...interface{}) {
	if !b.Enabled(context.Background(), slog.LevelWarn) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelWarn, fmt.Sprintf(format, args...), pcs[0])
	_ = b.Handler().Handle(context.Background(), r)
}

func (b Builder) Errorf(format string, args ...interface{}) {
	if !b.Enabled(context.Background(), slog.LevelError) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelError, fmt.Sprintf(format, args...), pcs[0])
	_ = b.Handler().Handle(context.Background(), r)
}

var globalOptions *Options
