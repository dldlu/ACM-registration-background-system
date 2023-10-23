package service

import "sync"

var fileService *FileService
var fileServiceOnce sync.Once

type FileService struct {
	ExportService
}

func GetFileService() *FileService {
	fileServiceOnce.Do(func() {
		fileService = &FileService{}
		exportService := new(ExportServiceImpl)

		fileService.ExportService = exportService
	})
	return fileService
}
