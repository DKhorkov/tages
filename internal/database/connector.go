package database

import (
	"database/sql"
	"log/slog"

	"github.com/DKhorkov/hmtm-sso/pkg/logging"
	"github.com/DKhorkov/tages/internal/config"
	customerrors "github.com/DKhorkov/tages/internal/errors"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

type CommonDBConnector struct {
	Connection *sql.DB
	Driver     string
	DSN        string
	Logger     *slog.Logger
}

func (connector *CommonDBConnector) Connect() error {
	if connector.Connection == nil {
		connection, err := sql.Open(connector.Driver, connector.DSN)

		if err != nil {
			return err
		}

		connector.Connection = connection
	}

	return nil
}

func (connector *CommonDBConnector) GetConnection() *sql.DB {
	if connector.Connection == nil {
		if err := connector.Connect(); err != nil {
			return nil
		}
	}

	return connector.Connection
}

func (connector *CommonDBConnector) GetTransaction() (*sql.Tx, error) {
	if connector.Connection == nil {
		return nil, &customerrors.NilDBConnectionError{}
	}

	return connector.Connection.Begin()
}

func (connector *CommonDBConnector) CloseConnection() {
	if connector.Connection == nil {
		return
	}

	if err := connector.Connection.Close(); err != nil {
		connector.Logger.Error(
			"Failed to close database connection",
			"Traceback",
			logging.GetLogTraceback(),
			"Error",
			err,
		)
	}
}

func New(dbConfig config.DatabaseConfig, logger *slog.Logger) (*CommonDBConnector, error) {
	dbConnector := &CommonDBConnector{
		Driver: dbConfig.Driver,
		DSN:    dbConfig.DSN,
		Logger: logger,
	}

	if err := dbConnector.Connect(); err != nil {
		return nil, err
	}

	return dbConnector, nil
}
