package utils

import "flag"

type Flags struct {
	Url    string
	Output string
}

func ParseFlags() Flags {
	var url, output string

	flag.StringVar(&url, "url", "", "The URL to download a file")
	flag.StringVar(&output, "output", "", "Name of the file")
	flag.Parse()

	return Flags{
		Url:    url,
		Output: output,
	}
}
