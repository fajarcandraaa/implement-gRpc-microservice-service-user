package userrepo

import (
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-user/helpers/errorcodehandling"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db        *gorm.DB
	codeError *errorcodehandling.CodeError
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
