package main

import (
	"bytes"
	"fmt"
	"io"

	"task/internal/readers"
	"task/pkg/variables"
)

func main() {
	buf := bytes.NewBufferString(variables.Str)
	limitedReader := readers.LimitReader(buf, int64(variables.BytesLength))

	result := make([]byte, 20)
	n, err := limitedReader.Read(result)
	if err != nil && err != io.EOF {
		panic(err)
	}

	fmt.Println(string(result[:n]))
}
