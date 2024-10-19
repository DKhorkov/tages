package storage

import (
	"bytes"
)

type File struct {
	Name      string
	Extension string
	Buffer    *bytes.Buffer
}

func (f *File) Write(chunk []byte) error {
	_, err := f.Buffer.Write(chunk)
	return err
}

func (f *File) Read(chunk []byte) (int, error) {
	return f.Buffer.Read(chunk)
}

func NewFile(name string, extension string) *File {
	return &File{
		Name:      name,
		Extension: extension,
		Buffer:    &bytes.Buffer{},
	}
}
