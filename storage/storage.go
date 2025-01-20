package storage

import "os"

type Storage interface {
	Mkdir(path string) error
	Write(path string, data []byte) error
	WriteStream(path string, file *os.File) error
	Delete(path string) error
}
