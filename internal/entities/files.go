package entities

import "time"

type FileMetadata struct {
	UUID      string
	Filename  string
	Extension string
	CreatedAt time.Time
	UpdatedAt time.Time
}
