package main

import (
	"github.com/sirupsen/logrus"
	"github.com/thorfour/coin/pkg/coin"
	"github.com/thorfour/sillyputty/pkg/sillyputty"
)

func main() {
	logrus.Info("Starting coin server")

	s := sillyputty.New("/v1",
		sillyputty.HandlerOpt("/coin", coin.SillyPuttyHandler),
	)
	s.Port = 4343
	s.Run()
}
