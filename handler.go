package pfxlog

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"runtime"
	"strings"
	"sync"
	"time"
)

type PrettyHandler struct {
	level   slog.Level
	options *Options
	lock    sync.Mutex
}

func NewPrettyHandler(level slog.Level, options *Options) slog.Handler {
	return &PrettyHandler{level: level, options: options}
}

func (h *PrettyHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *PrettyHandler) Handle(_ context.Context, r slog.Record) error {
	var timeLabel string
	if h.options.AbsoluteTime {
		timeLabel = "[" + time.Now().Format(h.options.TimestampFormat) + "]"
	} else {
		seconds := time.Since(h.options.StartTimestamp).Seconds()
		timeLabel = fmt.Sprintf("[%8.3f]", seconds)
	}

	var level string
	switch r.Level {
	case slog.LevelError:
		level = h.options.ErrorLabel
	case slog.LevelWarn:
		level = h.options.WarningLabel
	case slog.LevelInfo:
		level = h.options.InfoLabel
	case slog.LevelDebug:
		level = h.options.DebugLabel
	}

	fs := runtime.CallersFrames([]uintptr{r.PC})
	f, _ := fs.Next()
	functionStr := f.Function
	if h.options.TrimPrefix != "" {
		functionStr = strings.TrimPrefix(functionStr, h.options.TrimPrefix)
	}

	fieldsMap := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fieldsMap[a.Key] = a.Value.Any()
		return true
	})
	fieldsBytes, err := json.Marshal(fieldsMap)
	if err != nil {
		return err
	}
	fieldsStr := ""
	if len(fieldsBytes) > 2 {
		fieldsStr = h.options.FieldsColor + string(fieldsBytes) + h.options.DefaultFgColor
	}

	h.lock.Lock()
	fmt.Println(h.options.TimestampColor+timeLabel+h.options.DefaultFgColor,
		level,
		h.options.FunctionColor+functionStr+h.options.DefaultFgColor,
		fieldsStr,
		r.Message)
	h.lock.Unlock()

	return nil
}

func (h *PrettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *PrettyHandler) WithGroup(name string) slog.Handler {
	return h
}
