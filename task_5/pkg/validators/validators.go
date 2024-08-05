package validators

import (
	"errors"
	"net/url"
	"task/pkg/utils"
)

func ValidateFlags(flags utils.Flags) error {
	if flags.Url == "" {
		return errors.New("URL cannot be empty")
	}

	if _, err := url.ParseRequestURI(flags.Url); err != nil {
		return errors.New("invalid URL format")
	}

	if flags.Output == "" {
		return errors.New("output filename cannot be empty")
	}

	return nil
}
