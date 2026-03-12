package grpc_v1

import (
	"main/pkg"

	pb "github.com/cosmo-services/grpc-contracts/gen/api/v1"
	"go.uber.org/fx"
)

type GrpcHandler struct {
	grpc          pkg.GrpcServer
	socialHandler *SocialHandler
}

func NewGrpcHandler(
	grpc pkg.GrpcServer,
	socialHandler *SocialHandler,
) *GrpcHandler {
	return &GrpcHandler{
		grpc:          grpc,
		socialHandler: socialHandler,
	}
}

func (h *GrpcHandler) Setup() {
	pb.RegisterSocialServiceServer(
		h.grpc.Server,
		h.socialHandler,
	)
}

var Module = fx.Options(
	fx.Provide(NewGrpcHandler),
	fx.Provide(NewSocialHandler),
)
