package user

import (
	"context"

	proto "github.com/fajarcandraaa/implement-gRpc-microservice-protobank/grpc/pb"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-user/internal/entity/userentity"
)

type Service interface {
	InsertNewUser(ctx context.Context, payload *proto.CreateUserRequest) (*userentity.User, error)
	FindUser(ctx context.Context, userID string) (*userentity.User, error)
	// GetListUsers(ctx context.Context, sortBy, orderBy string, perPage, page int) (*[]userentity.Users, int64, error)
	// UpdateUser(ctx context.Context, payload *userentity.UserData) (*userentity.UserData, error)
	// DeleteDataUser(ctx context.Context, userID string) error
}
