package filestorage

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"strings"

	"github.com/DKhorkov/hmtm-sso/pkg/logging"
	"github.com/DKhorkov/tages/internal/entities"
	customerrors "github.com/DKhorkov/tages/internal/errors"
	"github.com/DKhorkov/tages/internal/interfaces"
	"github.com/DKhorkov/tages/internal/storage"
	filestorage "github.com/DKhorkov/tages/protobuf/generated/go/file_storage"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ServerAPI struct {
	// Helps to test single endpoints, if others is not implemented yet
	filestorage.UnimplementedFileServiceServer
	UseCases      interfaces.UseCases
	Logger        *slog.Logger
	ChunkSize     int
	StreamChannel chan int
	UnaryChannel  chan int
}

func (api *ServerAPI) Upload(stream filestorage.FileService_UploadServer) error {
	defer api.ReleaseStreamRequestCounter()

	var (
		file             *storage.File
		originalFilename string
		fileID           = uuid.New().String()
	)

	for {
		request, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			metadata := entities.FileMetadata{
				Filename:  originalFilename,
				Extension: file.Extension,
				UUID:      fileID,
			}

			if err = api.UseCases.UploadFile(file, metadata); err != nil {
				api.Logger.Error(
					"Error occurred while trying to upload file",
					"Traceback",
					logging.GetLogTraceback(),
					"Error",
					err,
				)

				return status.Error(codes.Internal, err.Error())
			}

			return stream.SendAndClose(&filestorage.UploadResponse{FileId: fileID})
		} else if err != nil {
			api.Logger.Error(
				"Error occurred while trying receive data from stream",
				"Traceback",
				logging.GetLogTraceback(),
				"Error",
				err,
			)

			return status.Error(codes.Internal, err.Error())
		}

		if file == nil {
			file = storage.NewFile(fileID, request.GetFileExtension())
		}

		if originalFilename == "" {
			split := strings.Split(request.GetFilename(), "/")
			originalFilename = split[len(split)-1] + "." + request.GetFileExtension()
		}

		if err = file.Write(request.GetChunk()); err != nil {
			api.Logger.Error(
				"Error occurred while trying write chunk to file",
				"Traceback",
				logging.GetLogTraceback(),
				"Error",
				err,
			)

			return status.Error(codes.Internal, err.Error())
		}
	}
}

func (api *ServerAPI) Download(
	request *filestorage.DownloadRequest,
	server filestorage.FileService_DownloadServer,
) error {
	defer api.ReleaseStreamRequestCounter()

	fileID := request.GetFileId()
	file, err := api.UseCases.DownloadFile(fileID)
	if err != nil {
		api.Logger.Error(
			"File was not found",
			"Traceback",
			logging.GetLogTraceback(),
			"Error",
			err,
		)

		return status.Error(codes.NotFound, customerrors.FileNotFoundError{}.Error())
	}

	responseChunk := &filestorage.DownloadResponse{
		Chunk:     make([]byte, api.ChunkSize),
		Filename:  file.Name,
		Extension: file.Extension,
	}

	for {
		readBytesNumber, err := file.Read(responseChunk.GetChunk())
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			api.Logger.Error(
				"Error occurred while trying read file",
				"Traceback",
				logging.GetLogTraceback(),
				"Error",
				err,
			)

			return status.Errorf(codes.Internal, err.Error())
		}

		responseChunk.Chunk = responseChunk.GetChunk()[:readBytesNumber]
		if err = server.Send(responseChunk); err != nil {
			api.Logger.Error(
				"Error occurred while trying send file chunk to client",
				"Traceback",
				logging.GetLogTraceback(),
				"Error",
				err,
			)

			return status.Errorf(codes.Internal, err.Error())
		}
	}

	return nil
}

func (api *ServerAPI) ShowFiles(ctx context.Context, request *emptypb.Empty) (*filestorage.ShowFilesResponse, error) {
	defer api.ReleaseUnaryRequestCounter()

	files, err := api.UseCases.ShowFiles()
	if err != nil {
		api.Logger.ErrorContext(
			ctx,
			"Error occurred while trying to show files",
			"Traceback",
			logging.GetLogTraceback(),
			"Error",
			err,
		)

		return nil, &customerrors.GRPCError{Status: codes.Internal, Message: err.Error()}
	}

	filesForResponse := make([]*filestorage.ShowFileResponse, len(files))
	for i, file := range files {
		filesForResponse[i] = &filestorage.ShowFileResponse{
			FileID:    file.UUID,
			Filename:  file.Filename,
			Extension: file.Extension,
			CreatedAt: timestamppb.New(file.CreatedAt),
			UpdatedAt: timestamppb.New(file.UpdatedAt),
		}
	}

	return &filestorage.ShowFilesResponse{Files: filesForResponse}, nil
}

func (api *ServerAPI) ReleaseStreamRequestCounter() {
	<-api.StreamChannel
}

func (api *ServerAPI) ReleaseUnaryRequestCounter() {
	<-api.UnaryChannel
}

// RegisterServer handler (serverAPI) for AuthServer  to gRPC server:.
func RegisterServer(
	gRPCServer *grpc.Server,
	useCases interfaces.UseCases,
	logger *slog.Logger,
	chunkSize int,
	streamChannel chan int,
	unaryChannel chan int,

) {
	filestorage.RegisterFileServiceServer(
		gRPCServer,
		&ServerAPI{
			UseCases:      useCases,
			Logger:        logger,
			ChunkSize:     chunkSize,
			StreamChannel: streamChannel,
			UnaryChannel:  unaryChannel,
		},
	)
}
