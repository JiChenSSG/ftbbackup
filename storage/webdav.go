package storage

var PEMESSION int = 0600

type Webdav struct {
}

// Mkdir implements Storage.
func (w *Webdav) Mkdir(path string) {
	panic("unimplemented")
}

// Write implements Storage.
func (w *Webdav) Write(path string, data []byte) {
	panic("unimplemented")
}



var _ Storage = (*Webdav)(nil)
