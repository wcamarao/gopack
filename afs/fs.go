// gopack - Golang asset pipeline
// https://github.com/wcamarao/gopack
// MIT Licensed
//
package gopackafs

import (
	"net/http"
)

// Composite FS to avoid directory listing.
// Used externally by http file servers.
//
type AssetFileSystem struct {
	Hfs http.FileSystem
}

// Compose Open function with http.File, wrapping the
// resulting file with assetFile.
//
func (afs AssetFileSystem) Open(name string) (http.File, error) {
	f, err := afs.Hfs.Open(name)
	if err != nil {
		return nil, err
	}
	return assetFile{f}, nil
}
