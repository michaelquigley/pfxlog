package other

import (
	"github.com/michaelquigley/pfxlog"
	"log/slog"
)

type Component struct{}

func (c *Component) Hello() {
	pfxlog.Debug("debugging")
	pfxlog.Infof("this is #%d", 6)
	pfxlog.Info("oh, wow!")
	slog.With("severity", "severe").Warn("oh, no!")
	pfxlog.Error("uh...")
}
