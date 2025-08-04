package handlers

import (
	"context"

	"github.com/valdevay/1-APIHandlers/internal/userService"
	"github.com/valdevay/1-APIHandlers/internal/web/users"
)

type UserHandler struct {
	service userService.UserService
}

func NewUserHandler(service userService.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	userList, err := h.service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	responseUsers := make([]users.User, 0)
	for _, user := range userList {
		responseUser := users.User{
			Id:       uint(user.ID),
			Email:    user.Email,
			Password: user.Password,
		}
		responseUsers = append(responseUsers, responseUser)
	}

	return users.GetUsers200JSONResponse(responseUsers), nil
}

func (h *UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	if request.Body == nil {
		return users.PostUsers201JSONResponse{}, nil
	}

	user := userService.User{
		Email:    request.Body.Email,
		Password: request.Body.Password,
	}

	createdUser, err := h.service.CreateUser(user)
	if err != nil {
		return nil, err
	}

	responseUser := users.User{
		Id:       uint(createdUser.ID),
		Email:    createdUser.Email,
		Password: createdUser.Password,
	}

	return users.PostUsers201JSONResponse(responseUser), nil
}

func (h *UserHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	if request.Body == nil {
		return users.PatchUsersId200JSONResponse{}, nil
	}

	user := userService.User{
		Email:    request.Body.Email,
		Password: request.Body.Password,
	}

	updatedUser, err := h.service.UpdateUser(int(request.Id), user)
	if err != nil {
		return nil, err
	}

	responseUser := users.User{
		Id:       uint(updatedUser.ID),
		Email:    updatedUser.Email,
		Password: updatedUser.Password,
	}

	return users.PatchUsersId200JSONResponse(responseUser), nil
}

func (h *UserHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	err := h.service.DeleteUser(int(request.Id))
	if err != nil {
		return nil, err
	}

	return users.DeleteUsersId204Response{}, nil
}
