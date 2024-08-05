package interfaces

type DownloadsInterface interface {
	DownloadFile(fileURL, fileName string) error
}
