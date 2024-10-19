package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"

	filestorage "github.com/DKhorkov/tages/protobuf/generated/go/file_storage"

	"github.com/DKhorkov/tages/internal/config"
	"github.com/DKhorkov/tages/internal/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	settings := config.New()
	clientConnection, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", settings.HTTP.Host, settings.HTTP.Port),
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)

	if err != nil {
		log.Fatalln(err)
	}

	defer clientConnection.Close()

	client := filestorage.NewFileServiceClient(clientConnection)

	var fileID string
	wg := new(sync.WaitGroup)
	for range 50 {
		wg.Add(1)
		time.Sleep(1 * time.Millisecond)
		go func() {
			fileID, err = upload(client, settings.Files.ChunkSize)
			if err != nil {
				log.Printf("Failed to upload file: too many connections")
			} else {
				log.Printf("Uploaded file id: %s\n", fileID)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	// Get number of files:
	response, err := client.ShowFiles(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Println(err)
	}

	log.Printf("Files len: %d", len(response.GetFiles()))

	// Download logics:
	err = download(client, fileID, settings.Files.DownloadDir)
	if err != nil {
		log.Println(err)
	}
}

func download(client filestorage.FileServiceClient, fileID string, downloadDir string) error {
	stream, err := client.Download(context.Background(), &filestorage.DownloadRequest{FileId: fileID})
	if err != nil {
		return err
	}

	var file *storage.File
	response := new(filestorage.DownloadResponse)
	for {
		err = stream.RecvMsg(response)
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return err
		}

		if file == nil {
			file = storage.NewFile(response.GetFilename(), response.GetExtension())
		}

		if len(response.GetChunk()) > 0 {
			err = file.Write(response.GetChunk())
			if err != nil {
				_ = stream.CloseSend()
				return err
			}
		}

		response.Chunk = response.GetChunk()[:0]
	}

	filePath := fmt.Sprintf("%s/%s-%s", downloadDir, fileID, file.Name)
	if err := os.WriteFile(filePath, file.Buffer.Bytes(), 0600); err != nil {
		return err
	}

	log.Printf("Downloaded file with id: %s. Path to file: %s", fileID, filePath)
	return nil
}

func upload(client filestorage.FileServiceClient, chunkSize int) (string, error) {
	stream, err := client.Upload(context.Background())
	if err != nil {
		return "", err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	filePath := cwd + "/assets/images/test.png" // filepath for tests
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return "", err
	}

	buffer := make([]byte, chunkSize)
	for {
		readBytesNumber, err := file.Read(buffer)
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return "", err
		}

		splitFilename := strings.Split(file.Name(), ".")
		request := &filestorage.UploadRequest{
			Chunk:         buffer[:readBytesNumber],
			Filename:      splitFilename[0],
			FileExtension: splitFilename[1],
		}

		if err = stream.Send(request); err != nil {
			return "", err
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		return "", err
	}

	return response.GetFileId(), nil
}
