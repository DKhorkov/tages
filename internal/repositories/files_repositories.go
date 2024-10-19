package repositories

import (
	"bytes"
	"fmt"
	"os"

	"github.com/DKhorkov/tages/internal/database"
	"github.com/DKhorkov/tages/internal/entities"
	"github.com/DKhorkov/tages/internal/interfaces"
	"github.com/DKhorkov/tages/internal/storage"
)

type CommonFilesStorageRepository struct {
	Dir string
}

func (repo *CommonFilesStorageRepository) UploadFile(file *storage.File) error {
	filePath := fmt.Sprintf("%s/%s.%s", repo.Dir, file.Name, file.Extension)
	if err := os.WriteFile(filePath, file.Buffer.Bytes(), 0600); err != nil {
		return err
	}

	return nil
}

func (repo *CommonFilesStorageRepository) DownloadFile(fileMetadata entities.FileMetadata) (*storage.File, error) {
	filePath := fmt.Sprintf("%s/%s.%s", repo.Dir, fileMetadata.UUID, fileMetadata.Extension)
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return &storage.File{
		Name:      fileMetadata.Filename,
		Extension: fileMetadata.Extension,
		Buffer:    bytes.NewBuffer(file),
	}, nil
}

type CommonFilesMetadataRepository struct {
	DBConnector interfaces.DBConnector
}

func (repo *CommonFilesMetadataRepository) SaveFileMetadata(metadata entities.FileMetadata) error {
	connection := repo.DBConnector.GetConnection()
	_, err := connection.Exec(
		`
			INSERT INTO files_metadata (UUID, filename, extension) 
			VALUES ($1, $2, $3)
		`,
		metadata.UUID,
		metadata.Filename,
		metadata.Extension,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repo *CommonFilesMetadataRepository) GetFileMetadataByID(fileID string) (*entities.FileMetadata, error) {
	file := &entities.FileMetadata{}
	columns := database.GetEntityColumns(file)
	connection := repo.DBConnector.GetConnection()
	err := connection.QueryRow(
		`
			SELECT * 
			FROM files_metadata AS f
			WHERE f.UUID = $1
		`,
		fileID,
	).Scan(columns...)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func (repo *CommonFilesMetadataRepository) ShowFiles() ([]*entities.FileMetadata, error) {
	connection := repo.DBConnector.GetConnection()
	rows, err := connection.Query(
		`
			SELECT * 
			FROM files_metadata
		`,
	)

	if err != nil {
		return nil, err
	}

	var files []*entities.FileMetadata
	for rows.Next() {
		file := &entities.FileMetadata{}
		columns := database.GetEntityColumns(file)
		err = rows.Scan(columns...)
		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}

	if err = rows.Close(); err != nil {
		return nil, err
	}

	return files, nil
}
