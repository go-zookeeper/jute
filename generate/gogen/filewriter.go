package gogen

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
)

type fileWriter struct {
	buf bytes.Buffer
}

func (w *fileWriter) printf(f string, args ...interface{}) {
	fmt.Fprintf(&w.buf, f, args...)
}

func (w *fileWriter) format() []byte {
	src, err := format.Source(w.buf.Bytes())
	if err != nil {
		panic(err.Error() + "\n" + w.buf.String())
	}
	return src
}

func (w *fileWriter) writeFile(filename string) error {
	if err := ioutil.WriteFile(filename, w.format(), 0644); err != nil {
		return fmt.Errorf("failed to write file '%s': %w", filename, err)
	}
	return nil
}
