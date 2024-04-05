package main

import (
	"fmt"
	"github.com/michaelquigley/pfxlog"
	"github.com/michaelquigley/pfxlog/examples/pfxlog-example-simple/other"
	"log/slog"
	"time"
)

func init() {
	pfxlog.GlobalInit(slog.LevelDebug, pfxlog.DefaultOptions().SetTrimPrefix("github.com/michaelquigley/"))
}

func main() {
	log := pfxlog.Logger()
	log.With(slog.String("hello", "world")).Info("hello world")

	notifications := make(chan int)
	for i := 0; i < 50; i++ {
		go counter(i, notifications)
	}

	for i := 0; i < 50; i++ {
		n := <-notifications
		slog.With(slog.Int("n", n)).Info("done")
	}

	slog.Info("complete")
}

func counter(number int, notify chan int) {
	log := pfxlog.ContextLogger(fmt.Sprintf("#%d", number))

	for i := 0; i < 5; i++ {
		pfxlog.Infof("visited %d.", i)
	}

	time.Sleep(1 * time.Second)

	c := &other.Component{}
	c.Hello()

	log.Info("complete")

	notify <- number
}
