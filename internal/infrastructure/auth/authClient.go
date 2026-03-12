package auth_infrastructure

import (
	"context"
	"fmt"
	"main/internal/config"
	"main/pkg"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"main/internal/domain/auth"

	pb "github.com/cosmo-services/grpc-contracts/gen/api/v1"
)

type AuthGrpcClient struct {
	client  pb.AuthServiceClient
	conn    *grpc.ClientConn
	timeout time.Duration
}

func NewAuthGrpcClient(grpcClient *pkg.GrpcClient, env config.Env) (auth.AuthClient, error) {
	conn, err := grpcClient.Connect(env.AuthServiceGrpcAddress)
	if err != nil {
		return nil, err
	}

	return &AuthGrpcClient{
		client:  pb.NewAuthServiceClient(conn),
		conn:    conn,
		timeout: 15 * time.Second,
	}, nil
}

func (c *AuthGrpcClient) GetUserById(userId string) (*auth.AuthUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)

	defer cancel()

	req := &pb.GetUserByIdRequest{
		UserId: userId,
	}

	resp, err := c.client.GetUserById(ctx, req)
	if err != nil {
		return nil, c.mapGRPCError(err)
	}

	return c.mapToDomainUser(resp), nil
}

func (c *AuthGrpcClient) GetUserByUsername(username string) (*auth.AuthUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)

	defer cancel()

	req := &pb.GetUserByUsernameRequest{
		Username: username,
	}

	resp, err := c.client.GetUserByUsername(ctx, req)
	if err != nil {
		return nil, c.mapGRPCError(err)
	}

	return c.mapToDomainUser(resp), nil
}

func (c *AuthGrpcClient) mapGRPCError(err error) error {
	st, ok := status.FromError(err)
	if !ok {
		return fmt.Errorf("unexpected error: %w", err)
	}

	switch st.Code() {
	case codes.NotFound:
		return auth.ErrUserNotFound
	case codes.DeadlineExceeded:
		return fmt.Errorf("request timeout: %w", err)
	case codes.Unavailable:
		return fmt.Errorf("service unavailable: %w", err)
	default:
		return fmt.Errorf("gRPC error (code=%s): %w", st.Code(), err)
	}
}

func (c *AuthGrpcClient) mapToDomainUser(response *pb.GetUserResponse) *auth.AuthUser {
	return &auth.AuthUser{
		ID:        response.User.Id,
		Email:     response.User.Email,
		Username:  response.User.Username,
		IsActive:  response.User.IsActive,
		CreatedAt: time.Unix(response.User.CreatedAt, 0),
	}
}
