package repositories

import (
	"context"

	"github.com/fajarcandraaa/implement-gRpc-microservice-service-user/internal/entity/userentity"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-user/repositories/userrepo"
	"github.com/jinzhu/gorm"
)

type User interface {
	SaveNewUser(ctx context.Context, payload *userentity.User) error
	FindUserByID(ctx context.Context, userID string) (*userentity.User, error)
}

func NewUser(db *gorm.DB) User {
	return userrepo.NewUserRepository(db)
}
