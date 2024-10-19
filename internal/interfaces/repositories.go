package interfaces

import (
	"github.com/DKhorkov/tages/internal/entities"
	"github.com/DKhorkov/tages/internal/storage"
)

type FilesStorageRepository interface {
	UploadFile(file *storage.File) error
	DownloadFile(fileMetadata entities.FileMetadata) (*storage.File, error)
}

type FilesMetadataRepository interface {
	SaveFileMetadata(metadata entities.FileMetadata) error
	GetFileMetadataByID(fileID string) (*entities.FileMetadata, error)
	ShowFiles() ([]*entities.FileMetadata, error)
}
