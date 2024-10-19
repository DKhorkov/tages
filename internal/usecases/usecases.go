package usecases

import (
	"github.com/DKhorkov/tages/internal/entities"
	"github.com/DKhorkov/tages/internal/interfaces"
	"github.com/DKhorkov/tages/internal/storage"
)

type CommonUseCases struct {
	FileService interfaces.FilesService
}

func (useCases *CommonUseCases) UploadFile(file *storage.File, metadata entities.FileMetadata) error {
	return useCases.FileService.UploadFile(file, metadata)
}

func (useCases *CommonUseCases) DownloadFile(fileID string) (*storage.File, error) {
	return useCases.FileService.DownloadFile(fileID)
}

func (useCases *CommonUseCases) ShowFiles() ([]*entities.FileMetadata, error) {
	return useCases.FileService.ShowFiles()
}
