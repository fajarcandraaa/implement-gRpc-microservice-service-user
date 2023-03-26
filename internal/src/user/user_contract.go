package user

import (
	"context"

	proto "github.com/fajarcandraaa/implement-gRpc-microservice-protobank/grpc/pb"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-user/internal/entity/userentity"
)

type Service interface {
	InsertNewUser(ctx context.Context, payload *proto.CreateUserRequest) (*userentity.User, error)
	FindUser(ctx context.Context, userID string) (*userentity.User, error)
}
