package handler

import (
	"context"

	"github.com/fajarcandraaa/implement-gRpc-microservice-protobank/grpc/pb"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-user/internal/src/user"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service user.Service
}

func NewUserHandler(service user.Service) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// RegisterNewUser is handler function to Handle user registration
func (uh *UserHandler) ServiceRegisterUser(ctx context.Context, payload *pb.CreateUserRequest) (*pb.UserStatusResponse, error) {

	_, err := uh.service.InsertNewUser(ctx, payload)
	if err != nil {

		return nil, err
	}

	response := &pb.UserStatusResponse{
		Status:  "Success",
		Message: "Success to register user",
	}

	return response, nil
}

// FindUserByUserID is handler function to Handle find user
func (ud *UserHandler) ServiceFindUserById(ctx context.Context, proto *pb.FindUserByIdRequest) (*pb.UserStatusResponse, error) {

	findUser, err := ud.service.FindUser(ctx, proto.Id)
	if err != nil {

		return nil, err

	}

	user := &pb.UserResponse{
		Id:        findUser.ID,
		Name:      findUser.Name,
		Email:     findUser.Email,
		Username:  findUser.Username,
		Password:  findUser.Password,
		CreatedAt: findUser.CreatedAt.String(),
		UpdatedAt: findUser.UpdatedAt.String(),
	}

	response := &pb.UserStatusResponse{
		Status:  "Success",
		Message: "User Found",
		Data:    user,
	}

	return response, nil
}
