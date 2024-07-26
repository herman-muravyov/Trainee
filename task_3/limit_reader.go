package main

import "io"

type limitedReader struct {
	reader io.Reader
	limit  int64
}

func LimitReader(r io.Reader, n int64) io.Reader {
	if n <= 0 {
		return &emptyReader{}
	}

	return &limitedReader{
		reader: r,
		limit:  n,
	}
}

func (lr *limitedReader) Read(p []byte) (int, error) {
	if lr.limit <= 0 {
		return 0, io.EOF
	}

	// If the length of "p" exceeds the remaining limit, it truncates the byte slice "p"
	// to ensure that only the remaining limit number of bytes are read from the underlying reader
	if int64(len(p)) > lr.limit {
		p = p[:lr.limit]
	}

	n, err := lr.reader.Read(p)
	// This snippet decrements the "lr.limit" by the number of bytes read "n" during the read operation
	lr.limit -= int64(n)

	if err == io.EOF && lr.limit > 0 {
		err = nil
	}

	return n, err
}
