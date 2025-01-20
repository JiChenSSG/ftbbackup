package main

import (
	"fmt"
	"time"

	"github.com/jichenssg/ftbbackup/config"
	"github.com/jichenssg/ftbbackup/logger"
	"github.com/jichenssg/ftbbackup/service"
	"github.com/jichenssg/ftbbackup/storage"
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

	entry, err := service.GetLatestFile(config.GetConfig().Location, "")
	if err != nil {
		zap.L().Error("get file error",
			zap.Error(err),
		)
		return
	}

	fileinfo, _ := entry.Info()

	zap.L().Info("get file success",
		zap.String("filename", entry.Name()),
		zap.String("size", fmt.Sprintf("%.2fmb", float64(fileinfo.Size())/1024/1024)),
	)

	// data, err := service.ReadFile(fmt.Sprintf("%s/%s", config.GetConfig().Location, file.Name()))
	// if err != nil {
	// 	zap.L().Fatal("read file error",
	// 		zap.Error(err),
	// 	)
	// }

	file, err := service.GetFile(fmt.Sprintf("%s/%s", config.GetConfig().Location, entry.Name()))
	if err != nil {
		zap.L().Fatal("get file error",
			zap.Error(err),
		)
	}

	zap.L().Info("uploading...")

	var s storage.Storage
	if config.GetConfig().Webdav {
		s = storage.GetWebdavStorage(
			config.GetConfig().WebdavRoot,
			config.GetConfig().WebdavUser,
			config.GetConfig().WebdavPassword,
			3,
		)

		if s != nil {
			zap.L().Info("webdav upload start")
			success := false
			for i := 0; i < 3; i++ {
				if i != 0 {
					zap.L().Info("retrying...")
				}

				err := service.Upload(s, config.GetConfig().WebdavStoragePath, entry.Name(), file)

				if err != nil {
					zap.L().Error("upload error",
						zap.Error(err),
					)

					time.Sleep(5 * time.Second)
					continue
				}

				success = true
				zap.L().Info("webdav upload success")
				break
			}

			if !success {
				s.Delete(fmt.Sprintf("%s/%s", config.GetConfig().WebdavStoragePath, entry.Name()))
				zap.L().Error("webdav upload failed")
			}
		}
	}

	zap.L().Info("upload finished")

}
