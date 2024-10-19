package main

import (
	"database/sql"
	"os"

	"github.com/DKhorkov/hmtm-sso/pkg/logging"
	"github.com/DKhorkov/tages/internal/app"
	"github.com/DKhorkov/tages/internal/config"
	grpccontroller "github.com/DKhorkov/tages/internal/controllers/grpc"
	"github.com/DKhorkov/tages/internal/database"
	"github.com/DKhorkov/tages/internal/repositories"
	"github.com/DKhorkov/tages/internal/services"
	"github.com/DKhorkov/tages/internal/usecases"
	"github.com/pressly/goose/v3"
)

func main() {
	settings := config.New()
	logger := logging.GetInstance(
		settings.Logging.Level,
		settings.Logging.LogFilePath,
	)

	dbConnector, err := database.New(
		settings.Database,
		logger,
	)

	defer dbConnector.CloseConnection()

	if err != nil {
		panic(err)
	}

	if err = lifespan(dbConnector.GetConnection(), &settings.Database); err != nil {
		panic(err)
	}

	filesMetadataRepository := &repositories.CommonFilesMetadataRepository{DBConnector: dbConnector}
	filesStorageRepository := &repositories.CommonFilesStorageRepository{Dir: settings.Files.UploadDir}
	filesService := &services.CommonFilesService{
		FilesMetadataRepository: filesMetadataRepository,
		FilesStorageRepository:  filesStorageRepository,
	}

	useCases := &usecases.CommonUseCases{FileService: filesService}
	controller := grpccontroller.New(
		settings.HTTP.Host,
		settings.HTTP.Port,
		settings.Files.ChunkSize,
		settings.RateLimits,
		useCases,
		logger,
	)

	application := app.New(controller)
	application.Run()
}

func lifespan(connection *sql.DB, dbConfig *config.DatabaseConfig) error {
	if err := goose.SetDialect(dbConfig.Driver); err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	err = goose.Up(
		connection,
		cwd+dbConfig.MigrationsDir,
	)

	if err != nil {
		return err
	}

	return nil
}
