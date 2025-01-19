package main

import (
	"github.com/jichenssg/ftbbackup/config"
	"github.com/jichenssg/ftbbackup/logger"
	"github.com/jichenssg/ftbbackup/service"
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

	file, err := service.GetLatestFile(config.GetConfig().Location, "")
	if err != nil {
		zap.L().Error("get file error",
			zap.Error(err),
		)
		return
	}

	fileinfo, _ := file.Info()

	zap.L().Info("get file success",
		zap.String("filename", file.Name()),
		zap.Int64("size", fileinfo.Size()),
	)

}
