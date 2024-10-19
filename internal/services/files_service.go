package services

import (
	"github.com/DKhorkov/tages/internal/entities"
	"github.com/DKhorkov/tages/internal/interfaces"
	"github.com/DKhorkov/tages/internal/storage"
)

type CommonFilesService struct {
	FilesStorageRepository  interfaces.FilesStorageRepository
	FilesMetadataRepository interfaces.FilesMetadataRepository
}

func (service *CommonFilesService) UploadFile(file *storage.File, metadata entities.FileMetadata) error {
	if err := service.FilesStorageRepository.UploadFile(file); err != nil {
		return err
	}

	return service.FilesMetadataRepository.SaveFileMetadata(metadata)
}

func (service *CommonFilesService) DownloadFile(fileID string) (*storage.File, error) {
	fileMetadata, err := service.FilesMetadataRepository.GetFileMetadataByID(fileID)
	if err != nil {
		return nil, err
	}

	return service.FilesStorageRepository.DownloadFile(*fileMetadata)
}

func (service *CommonFilesService) ShowFiles() ([]*entities.FileMetadata, error) {
	return service.FilesMetadataRepository.ShowFiles()
}
