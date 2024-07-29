package services

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"task/internal/services/helpers"
	"task/internal/services/interfaces"
	"task/pkg/errors"
	"task/pkg/logging"
)

type downloadsService struct {
	logger *logging.Logger
}

func NewDownloadsService(logger *logging.Logger) interfaces.DownloadsInterface {
	return &downloadsService{logger: logger}
}

func (s *downloadsService) DownloadFile(fileURL, fileName string) error {
	s.logger.Infof("Making a request for: %s", fileURL)

	resp, err := http.Get(fileURL)
	if err != nil {
		s.logger.Errorf("Error making GET request: %v", err)

		if helpers.IsConnectionError(err) {
			return errors.ErrConnectionFailed("Connection failed", err.Error())
		}

		return errors.ErrInvalidURL("Invalid URL", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		s.logger.Errorf("Non-OK HTTP status: %v", resp.StatusCode)

		if resp.StatusCode == http.StatusNotFound {
			return errors.ErrFileNotFound("File not found", fmt.Sprintf("Received status code: %d, - seems to be not found", resp.StatusCode))
		}

		return errors.ErrDownloadFailed("Failed to download file", fmt.Sprintf("Received status code: %d, - something went wrong", resp.StatusCode))
	}

	s.logger.Info("Creating output path...")

	file, err := os.Create(fileName)
	if err != nil {
		s.logger.Errorf("Error creating file: %v", err)
		return errors.ErrDownloadFailed("Failed to create file", err.Error())
	}
	defer file.Close()

	s.logger.Info("Coping a body from the response to the output file...")

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		s.logger.Errorf("Error writing to file: %v", err)
		return errors.ErrDownloadFailed("Failed to write to file", err.Error())
	}

	s.logger.Infof("File successfully downloaded to %s", fileName)
	return nil
}
