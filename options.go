package pfxlog

import (
	"fmt"
	"github.com/mgutz/ansi"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"
)

type Options struct {
	Handler        slog.Handler
	StartTimestamp time.Time
	AbsoluteTime   bool
	TrimPrefix     string

	ErrorLabel   string
	WarningLabel string
	InfoLabel    string
	DebugLabel   string

	TimestampColor string
	FunctionColor  string
	ChannelColor   string
	FieldsColor    string
	DefaultFgColor string

	TimestampFormat string
}

func DefaultOptions() *Options {
	options := &Options{
		StartTimestamp:  time.Now(),
		AbsoluteTime:    false,
		TimestampFormat: "2006-01-02 15:04:05.000",
	}

	if defaultEnv("PFXLOG_USE_COLOR", false) {
		return options.Color()
	} else {
		return options.NoColor()
	}
}

func (options *Options) SetHandler(h slog.Handler) *Options {
	options.Handler = h
	return options
}

func (options *Options) Starting(t time.Time) *Options {
	options.StartTimestamp = t
	return options
}

func (options *Options) StartingToday() *Options {
	now := time.Now()
	options.StartTimestamp = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return options
}

func (options *Options) SetAbsoluteTime() *Options {
	options.AbsoluteTime = true
	return options
}

func (options *Options) SetTrimPrefix(prefix string) *Options {
	options.TrimPrefix = prefix
	return options
}

func (options *Options) Color() *Options {
	options.ErrorLabel = ansi.Red + "  ERROR" + ansi.DefaultFG
	options.WarningLabel = ansi.Yellow + "WARNING" + ansi.DefaultFG
	options.InfoLabel = ansi.White + "   INFO" + ansi.DefaultFG
	options.DebugLabel = ansi.Blue + "  DEBUG" + ansi.DefaultFG

	options.TimestampColor = ansi.Blue
	options.FunctionColor = ansi.Cyan
	options.ChannelColor = ansi.Yellow
	options.FieldsColor = ansi.LightCyan
	options.DefaultFgColor = ansi.DefaultFG

	return options
}

func (options *Options) NoColor() *Options {
	options.ErrorLabel = "  ERROR"
	options.WarningLabel = "WARNING"
	options.InfoLabel = "   INFO"
	options.DebugLabel = "  DEBUG"

	options.TimestampColor = ""
	options.FunctionColor = ""
	options.ChannelColor = ""
	options.FieldsColor = ""
	options.DefaultFgColor = ""

	return options
}

func defaultEnv(env string, defaultValue bool) bool {
	if envStr := strings.ToLower(os.Getenv(env)); envStr != "" {
		if envValue, err := strconv.ParseBool(envStr); err == nil {
			return envValue
		} else {
			_, _ = fmt.Fprintf(os.Stderr, "error parsing environment variable '%s' (%v)\n", env, err)
		}
	}
	return defaultValue
}
