package helpers

import (
	"os"
	"task/pkg/errors"
	"task/pkg/logging"
)

func CreateOutputDirectory(outputDir string, logger *logging.Logger) error {
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		logger.Errorf("Error creating output directory: %v", err)
		return errors.ErrDownloadFailed("Failed to create output directory", err.Error())
	}
	logger.Infof("Output directory created: %s", outputDir)
	return nil
}
