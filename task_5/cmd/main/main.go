package main

import (
	"fmt"
	"task/internal/services"
	"task/internal/services/interfaces"
	"task/pkg/logging"
	"task/pkg/utils"
	"task/pkg/validators"
)

func initializeServices(logger *logging.Logger) interfaces.DownloadsInterface {
	downloadsSvc := services.NewDownloadsService(logger)

	return downloadsSvc
}

func main() {
	logger := logging.GetLogger()
	flags := utils.ParseFlags()
	validators.ValidateFlags(flags)

	logger.Info("Initializing services...")
	yandexSvc := initializeServices(logger)
	yandexSvc.DownloadFile(flags.Url, flags.Output)

	logger.Infof("File has been downloaded from the URL: %s", flags.Url)

	fmt.Printf("File is successful downloaded from the URL: %s\nOutput is here: outputs/%s", flags.Url, flags.Output)
}
