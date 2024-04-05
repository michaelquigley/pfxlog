package other

import (
	"github.com/michaelquigley/pfxlog"
)

type Component struct{}

func (c *Component) Hello() {
	pfxlog.Infof("this is #%d", 6)
	pfxlog.Logger().Info("oh, wow!")
	pfxlog.Logger().Error("uh...")
}
