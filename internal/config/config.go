package config

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/DKhorkov/hmtm-bff/pkg/loadenv"
	"github.com/DKhorkov/hmtm-sso/pkg/logging"
)

func New() *Config {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return &Config{
		HTTP: HTTPConfig{
			Host: loadenv.GetEnv("HOST", "0.0.0.0"),
			Port: loadenv.GetEnvAsInt("PORT", 8080),
		},
		Database: DatabaseConfig{
			Driver:        loadenv.GetEnv("DB_DRIVER", "sqlite3"),
			DSN:           loadenv.GetEnv("DB_DSN", cwd+"/tages.db"),
			MigrationsDir: loadenv.GetEnv("DB_MIGRATIONS_DIR", "/internal/database/migrations"),
		},
		Logging: LoggingConfig{
			Level:       logging.LogLevels.DEBUG,
			LogFilePath: fmt.Sprintf("logs/%s.log", time.Now().Format("02-01-2006")),
		},
		Files: FilesConfig{
			UploadDir:   loadenv.GetEnv("UPLOAD_DIR", cwd+"/upload"),
			DownloadDir: loadenv.GetEnv("DOWNLOAD_DIR", cwd+"/download"),
			ChunkSize:   loadenv.GetEnvAsInt("FILE_CHUNK_SIZE", 1024), // 1KB
		},
		RateLimits: RateLimitsConfig{
			Unary:  loadenv.GetEnvAsInt("UNARY_RATE_LIMIT", 100),
			Stream: loadenv.GetEnvAsInt("STREAM_RATE_LIMIT", 10),
		},
	}
}

type HTTPConfig struct {
	Host string
	Port int
}

type DatabaseConfig struct {
	Driver        string
	DSN           string
	MigrationsDir string
}

type LoggingConfig struct {
	Level       slog.Level
	LogFilePath string
}

type FilesConfig struct {
	UploadDir   string
	DownloadDir string
	ChunkSize   int
}

type RateLimitsConfig struct {
	Unary  int
	Stream int
}

type Config struct {
	HTTP       HTTPConfig
	Database   DatabaseConfig
	Logging    LoggingConfig
	Files      FilesConfig
	RateLimits RateLimitsConfig
}
