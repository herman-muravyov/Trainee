package main

import "io"

type emptyReader struct{}

func (er *emptyReader) Read(p []byte) (int, error) {
	return 0, io.EOF
}
