package pfxlog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"time"
)

type formatter struct {
	start   time.Time
	options *Options
}

func NewFormatter() logrus.Formatter {
	options := colorDefaultsOrNot()
	return &formatter{start: options.StartTimestamp, options: options}
}

func NewFormatterWithOptions(options *Options) logrus.Formatter {
	return &formatter{start: options.StartTimestamp, options: options}
}

func (f *formatter) Format(entry *logrus.Entry) ([]byte, error) {
	second := time.Since(f.start).Seconds()
	var level string
	switch entry.Level {
	case logrus.PanicLevel:
		level = f.options.PanicLabel
	case logrus.FatalLevel:
		level = f.options.FatalLabel
	case logrus.ErrorLevel:
		level = f.options.ErrorLabel
	case logrus.WarnLevel:
		level = f.options.WarningLabel
	case logrus.InfoLevel:
		level = f.options.InfoLabel
	case logrus.DebugLevel:
		level = f.options.DebugLabel
	case logrus.TraceLevel:
		level = f.options.TraceLabel
	}
	trimmedFunction := ""
	if entry.Caller != nil {
		trimmedFunction = strings.TrimPrefix(entry.Caller.Function, prefix)
	}
	if context, found := entry.Data["context"]; found {
		trimmedFunction += " [" + context.(string) + "]"
	}
	message := entry.Message
	if withFields(entry.Data) {
		fields := "{"
		field := 0
		for k, v := range entry.Data {
			if k != "context" {
				if field > 0 {
					fields += " "
				}
				field++
				fields += fmt.Sprintf("%s=[%v]", k, v)
			}
		}
		fields += "} "
		message = f.options.FieldsColor + fields + f.options.DefaultFgColor + message
	}
	return []byte(fmt.Sprintf("%s %s %s: %s\n",
			f.options.TimestampColor+fmt.Sprintf("[%8.3f]", second)+f.options.DefaultFgColor,
			level,
			f.options.FunctionColor+trimmedFunction+f.options.DefaultFgColor,
			message),
		),
		nil
}

func withFields(data map[string]interface{}) bool {
	if _, found := data["context"]; found {
		return len(data) > 1
	} else {
		return len(data) > 0
	}
}

func colorDefaultsOrNot() *Options {
	var options *Options
	if useColor, err := strconv.ParseBool(strings.ToLower(os.Getenv("PFXLOG_USE_COLOR"))); err == nil && useColor {
		options = DefaultOptions()
	} else {
		options = DefaultOptions().NoColor()
	}
	return options
}
