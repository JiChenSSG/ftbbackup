package storage

import (
	"os"
	"time"

	"github.com/studio-b12/gowebdav"
	"go.uber.org/zap"
)

var PEMESSION os.FileMode = 0600

type Webdav struct {
	client *gowebdav.Client
}

// WriteStream implements Storage.
func (w *Webdav) WriteStream(path string, file *os.File) error {
	return w.client.WriteStream(path, file, PEMESSION)
}

// Delete implements Storage.
func (w *Webdav) Delete(path string) error {
	return w.client.Remove(path)
}

// Mkdir implements Storage.
func (w *Webdav) Mkdir(path string) error {
	return w.client.MkdirAll(path, PEMESSION)
}

// Write implements Storage.
func (w *Webdav) Write(path string, data []byte) error {
	return w.client.Write(path, data, PEMESSION)
}

// GetWebdavStorage returns a new Storage.
// root: WebDAV Endpoint
// user: User
// password: Password
// retry: Retry times
// if connect failed, return nil
func GetWebdavStorage(root, user, password string, retry int) Storage {
	client := gowebdav.NewClient(root, user, password)
	success := false
	for i := 0; i < retry; i++ {
		err := client.Connect()
		if err != nil {
			zap.L().Error("webdav connect error", zap.Error(err))
			time.Sleep(30 * time.Second)
			continue
		}

		zap.L().Info("webdav connect success")
		success = true
		break
	}

	if !success {
		zap.L().Error("webdav connect failed")
		return nil
	}

	return &Webdav{client: client}
}

var _ Storage = (*Webdav)(nil)
