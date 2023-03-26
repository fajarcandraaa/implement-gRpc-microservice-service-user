package handler

import (
	"context"

	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	// "net/http"
	// "strconv"

	// "github.com/fajarcandraaa/implement-gRpc-microservice-service-user/helpers"
	// "github.com/fajarcandraaa/implement-gRpc-microservice-service-user/internal/entity"
	// "github.com/fajarcandraaa/implement-gRpc-microservice-service-user/internal/entity/userentity"
	"github.com/fajarcandraaa/implement-gRpc-microservice-protobank/grpc/pb"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-user/internal/src/user"
	// "github.com/gorilla/mux"
	// "github.com/pkg/errors"
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
	// responder := helpers.NewHTTPResponse("registerNewUser")

	_, err := uh.service.InsertNewUser(ctx, payload) //uh.service.InsertNewUser(&payload)
	if err != nil {
		// response := &pb.UserStatusResponse{
		// 	Status:  "Failed",
		// 	Message: fmt.Sprint(err),
		// }

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
		// response := &pb.UserStatusResponse{
		// 	Status:  "Not Found",
		// 	Message: fmt.Sprint(err),
		// 	Data:    nil,
		// }

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

// // GetUsers is handler function to Handle list of users
// func (uh *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
// 	var (
// 		param        = r.URL.Query()
// 		paramPage    = param.Get("page")
// 		paramPerPage = param.Get("per_page")
// 		paramOrderBy = param.Get("order_by")
// 		paramSortBy  = param.Get("sort_by")
// 		responder    = helpers.NewHTTPResponse("registerNewUser")
// 		ctx          = r.Context()
// 	)

// 	paginationParam, err := helpers.SetDefaultPginationParam(paramPage, paramPerPage, paramOrderBy, paramSortBy)
// 	if err != nil {
// 		responder.FieldErrors(w, err, http.StatusUnprocessableEntity, "value of query parameters has diferent type")
// 		return
// 	}
// 	sortBy := paginationParam.SortBy
// 	orderBy := paginationParam.OrderBy
// 	perPage := paginationParam.PerPage
// 	page, _ := strconv.Atoi(paginationParam.Page)

// 	users, total, err := uh.service.GetListUsers(ctx, sortBy, orderBy, int(perPage), page)
// 	if err != nil {
// 		causer := errors.Cause(err)
// 		switch causer {
// 		case entity.ErrUserNotExist:
// 			responder.ErrorJSON(w, http.StatusNotFound, "users list not found")
// 			return
// 		default:
// 			responder.FailureJSON(w, err, http.StatusInternalServerError)
// 			return
// 		}
// 	}

// 	pagination, err := helpers.GetPagination(helpers.PaginationParams{
// 		Path:        "list.users",
// 		Page:        strconv.Itoa(page),
// 		TotalRows:   int32(total),
// 		PerPage:     int32(perPage),
// 		OrderBy:     orderBy,
// 		SortBy:      sortBy,
// 		CurrentPage: int32(page),
// 	})
// 	if err != nil {
// 		responder.ErrorJSON(w, http.StatusConflict, "error pagination")
// 		return
// 	}

// 	responder.SuccessWithMeta(w, users, pagination, http.StatusOK, "uses list")
// 	return

// }

// // UpdateDataUsers is handler function to Handle update data from users
// func (uh *UserHandler) UpdateDataUsers(w http.ResponseWriter, r *http.Request) {
// 	var (
// 		responder = helpers.NewHTTPResponse("updateDataUser")
// 		ctx       = r.Context()
// 	)
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, fmt.Sprint(err))
// 		return
// 	}

// 	var payload userentity.UserData
// 	err = json.Unmarshal(body, &payload)
// 	if err != nil {
// 		responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, fmt.Sprint(err))
// 		return
// 	}

// 	updatedUser, err := uh.service.UpdateUser(ctx, &payload)
// 	if err != nil {
// 		causer := errors.Cause(err)
// 		switch causer {
// 		case entity.ErrUserNotExist:
// 			responder.FieldErrors(w, err, http.StatusNotExtended, err.Error())
// 			return
// 		default:
// 			responder.FieldErrors(w, err, http.StatusInternalServerError, fmt.Sprint(err))
// 			return
// 		}
// 	}
// 	responder.SuccessJSON(w, updatedUser, http.StatusCreated, "Succes to update data user")
// 	return
// }

// // UserDelete is handler function to Handle dalete (destroy) data from users
// func (uh *UserHandler) UserDelete(w http.ResponseWriter, r *http.Request) {
// 	var (
// 		responder = helpers.NewHTTPResponse("updateDataUser")
// 		ctx       = r.Context()
// 		userID    = mux.Vars(r)["id"]
// 	)

// 	err := uh.service.DeleteDataUser(ctx, userID)
// 	if err != nil {
// 		causer := errors.Cause(err)
// 		switch causer {
// 		case entity.ErrUserNotExist:
// 			responder.ErrorJSON(w, http.StatusNotFound, "user not found")
// 			return
// 		default:
// 			responder.FailureJSON(w, err, http.StatusInternalServerError)
// 			return
// 		}
// 	}

// 	responder.SuccessWithoutData(w, http.StatusOK, "successfully to delete user")
// 	return
// }
