package auth_grpc

import (
	"context"
	"fmt"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/general_settings/config"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/general_settings/database/models"
	authv1 "github.com/Garmonik/gRPC_chat/backend/protos/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"net"
	"strconv"
	"time"
)

type GRPCAuthClient struct {
	client authv1.AuthClient
	conn   *grpc.ClientConn
	log    *slog.Logger
}

func New(log *slog.Logger, cfg *config.Config) (*GRPCAuthClient, error) {
	host := cfg.GrpcAuth.Host
	if host == "" {
		host = "localhost"
	}

	port := cfg.GrpcAuth.Port
	if port == 0 {
		port = 44044
	}

	address := net.JoinHostPort(host, strconv.Itoa(port))

	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithIdleTimeout(cfg.GrpcAuth.Timeout),
		grpc.WithConnectParams(grpc.ConnectParams{
			MinConnectTimeout: cfg.GrpcAuth.Timeout,
		}),
	)
	if err != nil {
		log.Error("failed to create gRPC client", "address", address, "err", err)
		return nil, fmt.Errorf("failed to create gRPC client to %s: %w", address, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if !conn.WaitForStateChange(ctx, connectivity.Idle) {
		log.Warn("gRPC connection didn't initialize in time", "address", address)
	}

	client := authv1.NewAuthClient(conn)

	log.Info("gRPC auth client created", "address", address)
	return &GRPCAuthClient{
		client: client,
		conn:   conn,
		log:    log,
	}, nil
}

func (g *GRPCAuthClient) Close() error {
	return g.conn.Close()
}

func (g *GRPCAuthClient) Login(ctx context.Context, email, password, ipAddress string) (string, error) {
	resp, err := g.client.Login(ctx, &authv1.LoginRequest{
		Email:     email,
		Password:  password,
		IpAddress: ipAddress,
	})
	if err != nil {
		g.log.Error("gRPC Login failed", "error", err)
		return "", err
	}
	return resp.SessionUuid, nil
}

func (g *GRPCAuthClient) Register(ctx context.Context, email, password, name string) (int64, error) {
	resp, err := g.client.Register(ctx, &authv1.RegisterRequest{
		Email:    email,
		Password: password,
		Name:     name,
	})
	if err != nil {
		g.log.Error("gRPC Register failed", "error", err)
		return 0, err
	}
	return resp.UserId, nil
}

func (g *GRPCAuthClient) Logout(ctx context.Context, sessionUUID, userID string) (string, error) {
	resp, err := g.client.Logout(ctx, &authv1.LogoutRequest{
		SessionUuid: sessionUUID,
		UserId:      userID,
	})
	if err != nil {
		g.log.Error("gRPC Logout failed", "error", err)
		return "", err
	}
	return resp.Message, nil
}

func (g *GRPCAuthClient) SessionsList(ctx context.Context, userID uint64) ([]models.Session, error) {
	resp, err := g.client.GetSessions(ctx, &authv1.GetSessionsRequest{
		UserId: userID,
	})
	if err != nil {
		g.log.Error("gRPC Logout failed", "error", err)
		return nil, err
	}
	return resp.Sessions, nil
}
