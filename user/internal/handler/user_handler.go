package handler

import (
	"context"

	models "user-service/internal/model"
	"user-service/internal/service"

	userpb "github.com/khbdev/proto-online-test/proto/user"
)


type UserHandler struct {
	userpb.UnimplementedUserServiceServer
	service *service.UserService
}


func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{service: svc}
}


func (h *UserHandler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	user := &models.User{
		FirstName:       req.FirstName,
		LastName:        req.LastName,
		Phone:           req.Phone,
		Email:           req.Email,
		TgUsername:      req.TgUsername,
		Bolimlar:        req.Bolimlar,
		Savollar:        req.Savollar,
		Javoblar:        req.Javoblar,
		TogriJavoblar:   int(req.TogriJavoblar),
		NatogriJavoblar: int(req.NatogriJavoblar),
	}

	createdUser, err := h.service.Create(
		user.FirstName,
		user.LastName,
		user.Phone,
		user.Email,
		user.TgUsername,
		user.Bolimlar,
		user.Savollar,
		user.Javoblar,
		user.TogriJavoblar,
		user.NatogriJavoblar,
	)
	if err != nil {
		return nil, err
	}

	return &userpb.CreateUserResponse{
		User: toProto(createdUser),
	}, nil
}

func (h *UserHandler) GetAllUsers(ctx context.Context, req *userpb.GetAllUsersRequest) (*userpb.GetAllUsersResponse, error) {
	users, err := h.service.GetAll()
	if err != nil {
		return nil, err
	}

	var protoUsers []*userpb.User
	for _, u := range users {
		protoUsers = append(protoUsers, toProto(&u))
	}

	return &userpb.GetAllUsersResponse{
		Users: protoUsers,
	}, nil
}


func (h *UserHandler) GetUserByID(ctx context.Context, req *userpb.GetUserByIDRequest) (*userpb.GetUserByIDResponse, error) {
	user, err := h.service.GetByID(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &userpb.GetUserByIDResponse{
		User: toProto(user),
	}, nil
}


func (h *UserHandler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	update := &models.User{
		FirstName:       req.FirstName,
		LastName:        req.LastName,
		Phone:           req.Phone,
		Email:           req.Email,
		TgUsername:      req.TgUsername,
		Bolimlar:        req.Bolimlar,
		Savollar:        req.Savollar,
		Javoblar:        req.Javoblar,
		TogriJavoblar:   int(req.TogriJavoblar),
		NatogriJavoblar: int(req.NatogriJavoblar),
	}

	updatedUser, err := h.service.Update(uint(req.Id), update.FirstName, update.LastName, update.Phone, update.Email, update.TgUsername, update.Bolimlar, update.Savollar, update.Javoblar, update.TogriJavoblar, update.NatogriJavoblar)
	if err != nil {
		return nil, err
	}

	return &userpb.UpdateUserResponse{
		User: toProto(updatedUser),
	}, nil
}


func (h *UserHandler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	err := h.service.Delete(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &userpb.DeleteUserResponse{Success: true}, nil
}


func toProto(u *models.User) *userpb.User {
	if u == nil {
		return nil
	}
	return &userpb.User{
		Id:               uint64(u.ID),
		FirstName:        u.FirstName,
		LastName:         u.LastName,
		Phone:            u.Phone,
		Email:            u.Email,
		TgUsername:       u.TgUsername,
		Bolimlar:         u.Bolimlar,
		Savollar:         u.Savollar,
		Javoblar:         u.Javoblar,
		TogriJavoblar:    int32(u.TogriJavoblar),
		NatogriJavoblar:  int32(u.NatogriJavoblar),
		CreatedAt:        u.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:        u.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
