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

func Login(request models.Login) (*models.GenerateJWT, error) {
	dataUser, err := usersrepository.GetUser("", request.Email)
	if err != nil {
		return nil, err
	}

	jwt := models.GenerateJWT{
		UserId: dataUser.ID.String(),
		Name:   dataUser.Name,
		Email:  dataUser.Email,
		Role:   dataUser.Role,
	}

	tokenString, _, err := utilities.GenerateJWT(&jwt)
	if err != nil {
		return nil, err
	}

	_, errUserIsLogin := usersrepository.ChangeIsLogin(dataUser.ID.String(), tokenString)
	if errUserIsLogin != nil {
		return nil, errUserIsLogin
	}

	jwt.Token = tokenString

	return &jwt, nil
}
