package handlers

import (
	"context"

	"main.go/internal/userService"
	"main.go/internal/web/users"
)

type userHandler struct {
	Service *userService.UserService
}

func NewUsersHandler(service *userService.UserService) *userHandler {
	return &userHandler{Service: service}
}

func (h *userHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	getUsers, err := h.Service.GetUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range getUsers {
		user := users.User{
			UserId:     &usr.ID,
			Email:   &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}

	return response, nil
}

func (h userHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body

	userToCreate := userService.User{
		Email:   *userRequest.Email,
		Password: *userRequest.Password,
	}

	createdUser, err := h.Service.CreateUser(userToCreate)

	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		UserId:     &createdUser.ID,
		Email:   &createdUser.Email,
		Password: &createdUser.Password,
	}

	return response, nil
}

func (h *userHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	userID := request.Id
	userRequestUpdate := request.Body

	userToUpdate := userService.User{
		Email:   *userRequestUpdate.Email,
		Password: *userRequestUpdate.Password,
	}

	updatedUser, err := h.Service.PatchUser(userID, userToUpdate)
	if err != nil {
		return nil, err
	}

	response := users.PatchUsersId200JSONResponse{
		UserId:     &updatedUser.ID,
		Email:   &updatedUser.Email,
		Password: &updatedUser.Password,
	}

	return response, nil
}

func (h userHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	userID := request.Id

	err := h.Service.DeleteUser(userID)
	if err != nil {
		return nil, err
	}

	return users.DeleteUsersId204Response{}, nil
}