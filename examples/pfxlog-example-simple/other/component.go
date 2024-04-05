package other

import (
	"github.com/michaelquigley/pfxlog"
	"log/slog"
)

type Component struct{}

func (c *Component) Hello() {
	slog.Warn("this is #%d", 6)
	pfxlog.Logger().Info("oh, wow!")
	pfxlog.Logger().Error("uh...")
}
