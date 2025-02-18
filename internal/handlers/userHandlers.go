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

func (h *userHandler) GetUsers (_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	getUsers, err := h.Service.GetUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range getUsers {
		user := users.User{
			Id:     &usr.ID,
			Email:   &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}

	return response, nil
}

func (h *userHandler) PatchUsersById(ctx context.Context, request users.PatchUsersByIdRequestObject) (users.PatchUsersByIdResponseObject, error) {
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

	response := users.PatchUsersById200JSONResponse{
		Id:     &updatedUser.ID,
		Email:   &updatedUser.Email,
		Password: &updatedUser.Password,
	}

	return response, nil
}

func (h userHandler) PostUsersPost(ctx context.Context, request users.PostUsersPostRequestObject) (users.PostUsersPostResponseObject, error) {
	userRequest := request.Body

	userToCreate := userService.User{
		Email:   *userRequest.Email,
		Password: *userRequest.Password,
	}

	createdUser, err := h.Service.CreateUser(userToCreate)

	if err != nil {
		return nil, err
	}

	response := users.PostUsersPost201JSONResponse{
		Id:     &createdUser.ID,
		Email:   &createdUser.Email,
		Password: &createdUser.Password,
	}

	return response, nil
}

func (h userHandler) DeleteUsersDeleteId(ctx context.Context, request users.DeleteUsersDeleteIdRequestObject) (users.DeleteUsersDeleteIdResponseObject, error) {
	userID := request.Id

	err := h.Service.DeleteUser(userID)
	if err != nil {
		return nil, err
	}

	return users.DeleteUsersDeleteId204Response{}, nil
}