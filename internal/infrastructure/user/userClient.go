package user_infrastructure

import (
	"context"
	"fmt"
	"main/internal/config"
	"main/pkg"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	user_domain "main/internal/domain/user"

	pb "github.com/cosmo-services/grpc-contracts/gen/api/v1"
)

type UserGrpcClient struct {
	client  pb.AuthServiceClient
	conn    *grpc.ClientConn
	timeout time.Duration
}

func NewUserGrpcClient(grpcClient pkg.GrpcClient, env config.Env) (user_domain.UserClient, error) {
	conn, err := grpcClient.Connect(env.AuthServiceGrpcAddress)
	if err != nil {
		return nil, err
	}

	return &UserGrpcClient{
		client:  pb.NewAuthServiceClient(conn),
		conn:    conn,
		timeout: 15 * time.Second,
	}, nil
}

func (c *UserGrpcClient) GetUser(userId string) (*user_domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)

	defer cancel()

	req := &pb.GetUserRequest{
		UserId: userId,
	}

	resp, err := c.client.GetUser(ctx, req)
	if err != nil {
		return nil, c.mapGRPCError(err)
	}

	return c.mapToDomainUser(resp), nil
}

func (c *UserGrpcClient) mapGRPCError(err error) error {
	st, ok := status.FromError(err)
	if !ok {
		return fmt.Errorf("unexpected error: %w", err)
	}

	switch st.Code() {
	case codes.NotFound:
		return fmt.Errorf("user not found: %w", err)
	case codes.DeadlineExceeded:
		return fmt.Errorf("request timeout: %w", err)
	case codes.Unavailable:
		return fmt.Errorf("service unavailable: %w", err)
	default:
		return fmt.Errorf("gRPC error (code=%s): %w", st.Code(), err)
	}
}

func (c *UserGrpcClient) mapToDomainUser(response *pb.GetUserResponse) *user_domain.User {
	return &user_domain.User{
		ID:        response.User.Id,
		Email:     response.User.Email,
		Username:  response.User.Username,
		IsActive:  response.User.IsActive,
		CreatedAt: time.Unix(response.User.CreatedAt, 0),
	}
}
