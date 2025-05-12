package auth_grpc

import (
	"context"
	"fmt"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/general_settings/config"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/general_settings/database/models"
	authv1 "github.com/Garmonik/gRPC_chat/backend/protos/gen/go/auth"
	"github.com/google/uuid"
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
		log.Error("failed to create Auth gRPC client", "address", address, "err", err)
		return nil, fmt.Errorf("failed to create Auth gRPC client to %s: %w", address, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if !conn.WaitForStateChange(ctx, connectivity.Idle) {
		log.Warn("Auth gRPC connection didn't initialize in time", "address", address)
	}

	client := authv1.NewAuthClient(conn)

	log.Info("Auth gRPC auth client created", "address", address)
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
		g.log.Error("Auth gRPC Login failed", "error", err)
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
		g.log.Error("Auth gRPC Register failed", "error", err)
		return 0, err
	}
	return resp.UserId, nil
}

func (g *GRPCAuthClient) Logout(ctx context.Context, sessionUUID string, userID uint64) (string, error) {
	resp, err := g.client.Logout(ctx, &authv1.LogoutRequest{
		SessionUuid: sessionUUID,
		UserId:      userID,
	})
	if err != nil {
		g.log.Error("Auth gRPC Logout failed", "error", err)
		return "", err
	}
	return resp.Message, nil
}

func (g *GRPCAuthClient) SessionsList(ctx context.Context, userID uint64) ([]models.Session, error) {
	resp, err := g.client.GetSessions(ctx, &authv1.GetSessionsRequest{
		UserId: userID,
	})
	if err != nil {
		g.log.Error("Auth gRPC Logout failed", "error", err)
		return nil, err
	}
	sessions := make([]models.Session, 0, len(resp.Sessions))
	for _, s := range resp.Sessions {
		session, err := ConvertPBToSession(s)
		if err != nil {
			g.log.Error("Failed to convert session", "error", err)
			continue
		}
		sessions = append(sessions, session)
	}

	return sessions, nil
}

func ConvertPBToSession(pbSession *authv1.Session) (models.Session, error) {
	sessionID, err := uuid.Parse(pbSession.Id)
	if err != nil {
		return models.Session{}, fmt.Errorf("invalid session id: %w", err)
	}

	createdAt, err := time.Parse(time.RFC3339, pbSession.CreatedAt)
	if err != nil {
		return models.Session{}, fmt.Errorf("invalid created_at: %w", err)
	}

	expiresAt, err := time.Parse(time.RFC3339, pbSession.ExpiresAt)
	if err != nil {
		return models.Session{}, fmt.Errorf("invalid expires_at: %w", err)
	}

	return models.Session{
		ID: sessionID,
		User: models.User{
			ID:    uint(pbSession.User.Id),
			Name:  pbSession.User.Name,
			Email: pbSession.User.Email,
		},
		IPAddress: pbSession.IpAddress,
		CreatedAt: createdAt,
		ExpiresAt: expiresAt,
		IsClosed:  pbSession.IsClosed,
	}, nil
}
