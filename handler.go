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
	attrs   []slog.Attr
}

func NewPrettyHandler(level slog.Level, options *Options) slog.Handler {
	return &PrettyHandler{level: level, options: options}
}

func (h *PrettyHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *PrettyHandler) Handle(_ context.Context, r slog.Record) error {
	var out strings.Builder

	var timeLabel string
	if h.options.AbsoluteTime {
		timeLabel = "[" + time.Now().Format(h.options.TimestampFormat) + "]"
	} else {
		seconds := time.Since(h.options.StartTimestamp).Seconds()
		timeLabel = fmt.Sprintf("[%8.3f]", seconds)
	}
	out.WriteString(h.options.TimestampColor + timeLabel + h.options.DefaultFgColor)

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
	out.WriteString(" " + level)

	fs := runtime.CallersFrames([]uintptr{r.PC})
	f, _ := fs.Next()
	functionStr := f.Function
	if h.options.TrimPrefix != "" {
		functionStr = strings.TrimPrefix(functionStr, h.options.TrimPrefix)
	}
	out.WriteString(" " + h.options.FunctionColor + functionStr + h.options.DefaultFgColor)

	r.AddAttrs(h.attrs...)
	fieldsMap := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		if a.Key != ChannelKey {
			fieldsMap[a.Key] = a.Value.Any()
		} else {
			out.WriteString(h.options.ChannelColor + " |" + a.Value.String() + "|" + h.options.DefaultFgColor)
		}
		return true
	})
	fieldsBytes, err := json.Marshal(fieldsMap)
	if err != nil {
		return err
	}
	if len(fieldsBytes) > 2 {
		out.WriteString(" " + h.options.FieldsColor + string(fieldsBytes) + h.options.DefaultFgColor)
	}

	out.WriteString(" " + r.Message)

	h.lock.Lock()
	fmt.Println(out.String())
	h.lock.Unlock()

	return nil
}

func (h *PrettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &PrettyHandler{level: h.level, options: h.options, attrs: attrs}
}

func (h *PrettyHandler) WithGroup(name string) slog.Handler {
	return h
}
