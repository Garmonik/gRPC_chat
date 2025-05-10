package grpcapp

import (
	"fmt"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/grpc_auth"
	authinterfase "github.com/Garmonik/gRPC_chat/backend/auth/internal/pkg/interfase_lib"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
	host       string
}

func New(log *slog.Logger, authService authinterfase.Auth,
	port int, host string) *App {
	gRPCServer := grpc.NewServer()

	grpc_auth.RegisterServerAPI(gRPCServer, authService)
	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
		host:       host,
	}
}

func (app *App) MustRun() error {
	const op = "grpcapp.Run"
	log := app.log.With(
		slog.String("operation", op),
		slog.Int("port", app.port),
	)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", app.host, app.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("starting gRPC server", slog.String("address", lis.Addr().String()))
	if err := app.gRPCServer.Serve(lis); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (app *App) Stop() {
	const op = "grpcapp.Stop"

	app.log.With(slog.String("operation", op)).Info("stopping gRPC server", slog.Int("port", app.port))

	app.gRPCServer.GracefulStop()
}
