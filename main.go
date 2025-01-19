package main

import (
	"github.com/jichenssg/ftbbackup/logger"
	"go.uber.org/zap"
)

func main() {
	shutdown := make([]func(), 0)
	defer func() {
		for _, fn := range shutdown {
			fn()
		}
	}()

	shutdown = append(shutdown, logger.Init())

	zap.L().Info("init success")
}
