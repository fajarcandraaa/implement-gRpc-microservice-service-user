package routers

import (
	"github.com/fajarcandraaa/implement-gRpc-microservice-protobank/grpc/pb"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-user/handler"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-user/internal/src/user"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-user/repositories"
)

func (se *Serve) initializeRoutes() {
	//======================== REPOSITORIES ========================
	//initiate repository
	r := repositories.NewRepository(se.DB)

	//======================== ROUTER ========================
	//Setting Services
	//Setting User Service

	//=== USER ===
	userService := user.NewService(r)
	UserHandler := handler.NewUserHandler(userService)

	pb.RegisterUserServiceServer(grpcServer, UserHandler)
	//=========================================================

	//======================== ENDPOINT ========================
	//Initialize endpoint route

	//=== USER ===
	// se.Router.HandleFunc("/user/create", h.RegisterNewUser).Methods("POST")
	// se.Router.HandleFunc("/user/{id}/find", h.FindUserByUserID).Methods("GET")
	// se.Router.HandleFunc("/user", h.GetUsers).Methods("GET")
	// se.Router.HandleFunc("/user", h.UpdateDataUsers).Methods("PUT")
	// se.Router.HandleFunc("/user/{id}", h.UserDelete).Methods("DELETE")
	//==========================================================

}
