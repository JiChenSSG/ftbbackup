package service

import (
	"fmt"
	"os"

	"github.com/jichenssg/ftbbackup/storage"
)

func Upload(s storage.Storage, path string, name string, file *os.File) error {
	if err := s.Mkdir(path); err != nil {
		return err
	}

	if err := s.WriteStream(fmt.Sprintf("%s/%s", path, name), file); err != nil {
		return err
	}

	return nil
}