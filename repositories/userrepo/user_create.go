package userrepo

import (
	"context"

	"github.com/pkg/errors"

	"github.com/fajarcandraaa/implement-gRpc-microservice-service-user/config/app"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-user/internal/entity"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-user/internal/entity/userentity"
)

// SaveNewUser is used to run query insert
func (r *UserRepository) SaveNewUser(ctx context.Context, payload *userentity.User) error {

	err := r.db.Create(payload).Error
	if err != nil {
		parsed := r.codeError.ParseSQLError(err)
		switch parsed {
		case app.ErrNoRowsFound:
			return entity.ErrUserNotExist
		case app.ErrUniqueViolation:
			return entity.ErrUserAlreadyExist
		default:
			return errors.Wrap(parsed, "build statement query to insert user from database")
		}
	}
	return nil
}
