// gopack - Golang asset pipeline
// https://github.com/wcamarao/gopack
// MIT Licensed
//
package gopack

// Interface for asset processors, allowing
// processors to be piped into one another.
//
type Processor interface {
	Process(filename string, bytes []byte) (string, []byte)
}
