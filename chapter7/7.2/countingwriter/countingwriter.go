package countingwriter

import (
	"io"
)

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &countingWriter{w: w}

	return cw, &cw.written
}

type countingWriter struct {
	w       io.Writer
	written int64
}

func (cw *countingWriter) Write(p []byte) (int, error) {
	n, err := cw.w.Write(p)
	if err == nil {
		cw.written += int64(n)
	}

	return n, err
}
