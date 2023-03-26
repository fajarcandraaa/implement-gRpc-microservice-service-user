package user

import (
	"context"
	"time"

	proto "github.com/fajarcandraaa/implement-gRpc-microservice-protobank/grpc/pb"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-user/helpers"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-user/helpers/errorcodehandling"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-user/helpers/unique"

	"github.com/fajarcandraaa/implement-gRpc-microservice-service-user/internal/entity"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-user/internal/entity/userentity"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-user/repositories"
	"github.com/google/uuid"
)

type service struct {
	repo *repositories.Repository
	err  *errorcodehandling.CodeError
}

func NewService(repo *repositories.Repository) *service {
	return &service{
		repo: repo,
	}
}

// InsertNewUser represents algorithm to register new user
func (s *service) InsertNewUser(ctx context.Context, payload *proto.CreateUserRequest) (*userentity.User, error) {

	err := userentity.UserRequestValidate(payload)
	if err != nil {
		return nil, err
	}
	hashPassword, _ := helpers.HashPassword(payload.Password)

	user := &userentity.User{
		ID:        uuid.NewString(),
		Name:      payload.Name,
		Email:     payload.Email,
		Username:  payload.Username,
		Password:  hashPassword,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	err = s.repo.User.SaveNewUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindUser represents algorithm to find user by id
func (s *service) FindUser(ctx context.Context, userID string) (*userentity.User, error) {
	if err := unique.ValidateUUID(userID); err != nil {
		return nil, entity.ErrUserNotExist
	}

	user, err := s.repo.User.FindUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
