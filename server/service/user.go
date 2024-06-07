package service

import (
	"net/rpc"
	"server/models"
)

type UserService struct {
	Client *rpc.Client
}

func NewUserService(client *rpc.Client) *UserService {
	return &UserService{Client: client}
}

func (s *UserService) CreateUser(user models.UserRequest) (*models.UserResponse, error) {
	var resp models.UserResponse
	err := s.Client.Call("UserService.CreateUser", user, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *UserService) GetAllUsers() ([]*models.UserResponse, error) {
	var resp []*models.UserResponse
	err := s.Client.Call("UserService.GetAllUsers", "", &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *UserService) GetUserByID(id int) (*models.UserResponse, error) {
	var resp models.UserResponse
	err := s.Client.Call("UserService.GetUserByID", id, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *UserService) UpdateUserByID(id int, user models.UserRequest) (*models.UserResponse, error) {
	var resp models.UserResponse
	err := s.Client.Call("UserService.UpdateUserByID", user, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *UserService) DeleteUserByID(id int) (*models.UserResponse, error) {
	var resp models.UserResponse
	err := s.Client.Call("UserService.DeleteUserByID", id, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
