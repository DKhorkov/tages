package grpccontroller

import (
	"fmt"

	"github.com/DKhorkov/tages/internal/controllers/ratelimiters"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/ratelimit"

	"github.com/DKhorkov/tages/internal/config"

	// ratelimit "github.com/tommy-sho/rate-limiter-grpc-go".

	"log/slog"
	"net"

	filestorage "github.com/DKhorkov/tages/internal/controllers/grpc/file_storage"

	"github.com/DKhorkov/tages/internal/interfaces"

	"github.com/DKhorkov/hmtm-sso/pkg/logging"
	"google.golang.org/grpc"
)

type Controller struct {
	grpcServer *grpc.Server
	host       string
	port       int
	logger     *slog.Logger
}

// Run gRPC server.
func (controller *Controller) Run() {
	controller.logger.Info(
		fmt.Sprintf("Starting gRPC Server at http://%s:%d", controller.host, controller.port),
		"Traceback",
		logging.GetLogTraceback(),
	)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", controller.host, controller.port))
	if err != nil {
		controller.logger.Error(
			"Failed to start gRPC Server",
			"Traceback",
			logging.GetLogTraceback(),
			"Error",
			err,
		)
		panic(err)
	}

	if err = controller.grpcServer.Serve(listener); err != nil {
		controller.logger.Error(
			"Error occurred while listening to gRPC server",
			"Traceback",
			logging.GetLogTraceback(),
			"Error",
			err,
		)
		panic(err)
	}

	controller.logger.Info("Stopped serving new connections.")
}

// Stop gRPC server gracefully (graceful shutdown).
func (controller *Controller) Stop() {
	// Stops accepting new requests and processes already received requests:
	controller.grpcServer.GracefulStop()
	controller.logger.Info("Graceful shutdown completed.")
}

// New creates an instance of gRPC Controller.
func New(
	host string,
	port int,
	chunkSize int,
	rateLimits config.RateLimitsConfig,
	useCases interfaces.UseCases,
	logger *slog.Logger,
) *Controller {
	streamChannel := make(chan int, rateLimits.Stream)
	unaryChannel := make(chan int, rateLimits.Unary)

	grpcServer := grpc.NewServer(
		grpc.ChainStreamInterceptor(
			ratelimit.StreamServerInterceptor(ratelimiters.NewLimiter(rateLimits.Stream, streamChannel)),
		),
		grpc.ChainUnaryInterceptor(
			ratelimit.UnaryServerInterceptor(ratelimiters.NewLimiter(rateLimits.Unary, unaryChannel)),
		),
	)

	// Connects our gRPC services to grpcServer:
	filestorage.RegisterServer(grpcServer, useCases, logger, chunkSize, streamChannel, unaryChannel)

	return &Controller{
		grpcServer: grpcServer,
		port:       port,
		host:       host,
		logger:     logger,
	}
}
