package limitreader

import (
	"io"
)

type limReader struct {
	r         io.Reader
	remaining int64
}

func (l *limReader) Read(p []byte) (int, error) {
	if l.remaining <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > l.remaining {
		p = p[0:l.remaining]
	}

	n, err := l.r.Read(p)
	l.remaining -= int64(n)

	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limReader{
		r:         r,
		remaining: n,
	}
}
