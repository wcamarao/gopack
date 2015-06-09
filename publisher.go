// gopack - Golang asset pipeline
// https://github.com/wcamarao/gopack
// MIT Licensed
//
package gopack

// Interface for asset publishers, allowing assets to be
// post-processed for tasks such as concat and minify.
//
// assets - Map of asset sources, where the key is each
//          filepath, and the value is an array of bytes,
//          representing the corresponding file contents.
//
// Return an array of published filename(s), or error.
//
type Publisher interface {
	Publish(assets map[string][]byte) ([]string, error)
}
