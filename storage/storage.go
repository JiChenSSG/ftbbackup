package storage

type Storage interface {
	Mkdir(path string)
	Write(path string, data []byte)
}