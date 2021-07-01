package pfxlog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strconv"
	"strings"
)

func GlobalInit(level logrus.Level, options *Options) {
	noJsonStr := "false"
	if noJsonEnv := strings.ToLower(os.Getenv("PFXLOG_NO_JSON")); noJsonEnv != "" {
		noJsonStr = noJsonEnv
	}
	noJson, err := strconv.ParseBool(noJsonStr)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "pfxlog: error parsing 'PFXLOG_NO_JSON' (%v)\n", err)
	}
	if (err == nil && noJson) || terminal.IsTerminal(int(os.Stdout.Fd())) {
		logrus.SetFormatter(NewFormatter(options))
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: options.JsonTimestampFormat})
	}
	logrus.SetLevel(level)
	logrus.SetReportCaller(true)
	globalOptions = options
	noLogger = logrus.New()
	noLogger.SetLevel(logrus.PanicLevel)
}

func Logger() *logrus.Entry {
	return logrus.NewEntry(logrus.StandardLogger())
}

func ContextLogger(context string) *logrus.Entry {
	return logrus.StandardLogger().WithField("context", context)
}

func ContextDataLogger(contextData interface{}) *logrus.Entry {
	if globalOptions.ContextDataFielder != nil {
		return globalOptions.ContextDataFielder(contextData, logrus.StandardLogger())
	} else {
		return logrus.StandardLogger().WithFields(nil)
	}
}

func ContextCheck(contextData interface{}) *logrus.Entry {
	if globalOptions.ContextChecker != nil && globalOptions.ContextChecker(contextData) {
		return logrus.StandardLogger().WithFields(nil)
	} else {
		return &logrus.Entry{Logger: noLogger}
	}
}

func ContextCheckData(contextData interface{}) *logrus.Entry {
	if globalOptions.ContextChecker != nil && globalOptions.ContextChecker(contextData) {
		return ContextDataLogger(contextData)
	} else {
		return &logrus.Entry{Logger: noLogger}
	}
}

var globalOptions *Options
var noLogger *logrus.Logger
