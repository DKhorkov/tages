package interfaces

import (
	"github.com/DKhorkov/tages/internal/entities"
	"github.com/DKhorkov/tages/internal/storage"
)

type UseCases interface {
	UploadFile(file *storage.File, metadata entities.FileMetadata) error
	DownloadFile(fileID string) (*storage.File, error)
	ShowFiles() ([]*entities.FileMetadata, error)
}
