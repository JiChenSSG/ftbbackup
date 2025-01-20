package storage_test

import (
	"os"
	"testing"

	"github.com/jichenssg/ftbbackup/config"
	"github.com/jichenssg/ftbbackup/storage"
)

var s storage.Storage

func init() {
	if err := os.Chdir("../"); err != nil {
		panic("Failed to change working directory")
	}

	s = storage.GetWebdavStorage(
		config.GetConfig().WebdavRoot,
		config.GetConfig().WebdavUser,
		config.GetConfig().WebdavPassword,
		3,
	)

	if s == nil {
		panic("webdav connect failed")
	}
}

func TestWebdavMkdir(t *testing.T) {
	if err := s.Mkdir("/public/test"); err != nil {
		t.Fatalf("Mkdir failed: %v", err)
	}
}

func TestWebdavWrite(t *testing.T) {
	if err := s.Mkdir("/public/test"); err != nil {
		t.Fatalf("Mkdir failed: %v", err)
	}

	if err := s.Write("/public/test/test.txt", []byte("test")); err != nil {
		t.Fatalf("Write failed: %v", err)
	}
}
