package authservice

import (
	"online-store/models"
	usersrepository "online-store/repositories/users_repository"
	"online-store/utilities"
)

func RegisterCustomer(request models.User) (*models.User, error) {
	request.Token, _ = utilities.HashPassword(request.Password)
	request.Role = "Customer"
	createUser, err := usersrepository.CreateUser(request)
	if err != nil {
		return nil, err
	}

	return createUser, nil
}

func RegisterAdmin(request models.User) (*models.User, error) {
	request.Token, _ = utilities.HashPassword(request.Password)
	request.Role = "Admin"
	createUser, err := usersrepository.CreateUser(request)
	if err != nil {
		return nil, err
	}

	return createUser, nil
}
