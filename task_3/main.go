package main

import (
	"bytes"
	"fmt"
	"io"

	"task/variables"
)

func main() {
	buf := bytes.NewBufferString(variables.Str)
	limitedReader := LimitReader(buf, int64(variables.BytesLength))

	result := make([]byte, 20)
	n, err := limitedReader.Read(result)
	if err != nil && err != io.EOF {
		panic(err)
	}

	fmt.Println(string(result[:n]))
}
