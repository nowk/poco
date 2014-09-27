package poco

import "io"
import "github.com/nowk/poc"

// Pipe returns a r, w both linked to a single poc. Follows the io.Pipe API
func Pipe() (io.Reader, io.Writer) {
	p := poc.New()
	r := &PocReader{
		poc: p,
	}
	w := &PocWriter{
		poc: p,
	}

	return r, w
}

// PocReader is the read interface of a single poc
type PocReader struct {
	poc *poc.Poc
}

// Read calls poc.Read(b)
func (p PocReader) Read(b []byte) (int, error) {
	return p.poc.Read(b)
}

// PocWriter is the write interface of a single poc
type PocWriter struct {
	poc *poc.Poc
}

// Write calls poc.Write(b)
func (p PocWriter) Write(b []byte) (int, error) {
	return p.poc.Write(b)
}
