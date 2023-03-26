package userrepo

import (
	"context"

	"github.com/fajarcandraaa/implement-gRpc-microservice-service-user/config/app"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-user/internal/entity/userentity"
	"github.com/pkg/errors"
)

// FindUserByID is used to run query find user
func (r *UserRepository) FindUserByID(ctx context.Context, userID string) (*userentity.User, error) {
	var user userentity.User
	err := r.db.First(&user, "id = ?", userID).Error
	if err != nil {
		parsed := r.codeError.ParseSQLError(err)
		switch parsed {
		case app.ErrNoRowsFound:
			return nil, app.ErrNoRowsFound
		case app.ErrUniqueViolation:
			return nil, app.ErrUniqueViolation
		default:
			return nil, errors.Wrap(parsed, "build statement query to find user from database")
		}
	}
	return &user, nil
}
