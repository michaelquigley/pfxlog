package pfxlog

import (
	"golang.org/x/crypto/ssh/terminal"
	"log/slog"
	"os"
)

func init() {
	// cover cases where ContextLogger is used in a package init function.
}

func GlobalInit(level slog.Level, options *Options) {
	if defaultEnv("PFXLOG_NO_JSON", false) || terminal.IsTerminal(int(os.Stdout.Fd())) {
		// pretty
	} else {
		// json
	}
	//logrus.SetLevel(level)
	//logrus.SetReportCaller(true)

	globalOptions = options
}

func GlobalConfig(f func(*Options) *Options) {
	globalOptions = f(globalOptions)
}

func Logger() *slog.Logger {
	//return &Builder{Entry: logrus.NewEntry(globalOptions.StandardLogger)}
	return nil
}

func ContextLogger(context string) *slog.Logger {
	//return &Builder{Entry: globalOptions.StandardLogger.WithField("_context", context)}
	return nil
}

type Builder struct {
	channels []string
	record   *slog.Record
}

func (self *Builder) Channels(channels ...string) *Builder {
	return self
}

func (self *Builder) WithChannels(channels ...string) *Builder {
	self.record.AddAttrs(slog.Any("_channels", self.channels))
	return self
}

func (self *Builder) SetContext(context string) *Builder {
	self.record.AddAttrs(slog.String("_context", context))
	return self
}

var globalOptions *Options
