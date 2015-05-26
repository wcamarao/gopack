// gopack - Golang asset pipeline
// https://github.com/wcamarao/gopack
// MIT Licensed
//
package gopackafs

import (
	"net/http"
	"os"
)

// Composite File to avoid directory listing.
// Used internally by the AssetFileSystem.
//
type assetFile struct {
	http.File
}

// Define Readdir to return nothing.
//
func (af assetFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}
