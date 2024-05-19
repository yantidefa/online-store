package authservice

import (
	"errors"
	"online-store/constants"
	"online-store/models"
	usersrepository "online-store/repositories/users_repository"
	"online-store/utilities"
)

func RegisterCustomer(request models.User) (*models.User, error) {
	request.Token, _ = utilities.HashPassword(request.Password)
	request.Role = "Customer"
	var createUser *models.User
	_, err := usersrepository.GetUser("", request.Email)
	if err != nil {
		createUser, err = usersrepository.CreateUser(request)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New(constants.ErrExistEmail)
	}

	return createUser, nil
}

func RegisterAdmin(request models.User) (*models.User, error) {
	dataUser, err := usersrepository.GetUser("", request.Email)
	if err != nil {
		return nil, err
	}
	if dataUser.Email == request.Email {
		return nil, errors.New(constants.ErrExistEmail)
	}

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

	if dataUser.Password != request.Password {
		return nil, errors.New(constants.ErrPassword)
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

	_, errUserIsLogin := usersrepository.ChangeIsLogin(dataUser.ID.String(), tokenString, true)
	if errUserIsLogin != nil {
		return nil, errUserIsLogin
	}

	jwt.Token = tokenString

	return &jwt, nil
}

func Logout(request models.Login) (int64, error) {
	dataUser, err := usersrepository.GetUser("", request.Email)
	if err != nil {
		return 0, err
	}

	userIsLogin, errUserIsLogin := usersrepository.ChangeIsLogin(dataUser.ID.String(), dataUser.Token, true)
	if errUserIsLogin != nil {
		return userIsLogin, errUserIsLogin
	}

	return userIsLogin, nil
}
