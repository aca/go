package iox

import "io"

type funcCloser struct {
	io.Reader
	close func() error
}

func (r funcCloser) Close() error {
	return r.close()
}

func FuncCloser(r io.Reader, close func() error) io.ReadCloser {
	return funcCloser{
		Reader: r,
		close:  close,
	}
}
