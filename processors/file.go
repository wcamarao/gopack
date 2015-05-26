// gopack - Golang asset pipeline
// https://github.com/wcamarao/gopack
// MIT Licensed
//
package processors

// File Processor
//
type FileProcessor struct {
}

// Keep original file name and contents.
//
func (fp FileProcessor) Process(filename string, bytes []byte) (string, []byte) {
	return filename, bytes
}
