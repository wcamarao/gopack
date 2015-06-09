// gopack - Golang asset pipeline
// https://github.com/wcamarao/gopack
// MIT Licensed
//
package processors

import (
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

// Sass Processor
//
type SassProcessor struct {
}

// Rename file to .css and compile source using sassc. If compilation fails,
// the output file will contain the error message from the sass compiler.
//
func (sp SassProcessor) Process(filename string, bytes []byte) (string, []byte) {
	noext := strings.TrimSuffix(filename, filepath.Ext(filename))
	newname := noext + ".css"

	cmd := exec.Command("sassc", "--stdin")
	stdin, stdout, stderr := sp.getStreams(cmd)

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	stdin.Write(bytes)
	stdin.Close()

	outbytes, errbytes := sp.readBytes(stdout), sp.readBytes(stderr)

	cmd.Wait()

	if len(outbytes) > 0 {
		return newname, outbytes
	} else {
		return newname, errbytes
	}
}

// Get all standard channels to/from a Cmd.
//
// Return (stdin, stdout, stderr)
//
func (sp SassProcessor) getStreams(cmd *exec.Cmd) (stdin io.WriteCloser, stdout, stderr io.ReadCloser) {
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	stdout, err = cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	stderr, err = cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	return stdin, stdout, stderr
}

// Read all bytes from a ReadCloser.
//
// Return array of bytes.
//
func (sp SassProcessor) readBytes(rc io.ReadCloser) []byte {
	bytes, err := ioutil.ReadAll(rc)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}
