package usersrepository

import (
	"online-store/config"
	"online-store/models"
	"online-store/utilities"
)

func CreateUser(request models.User) (*models.User, error) {
	request.CreatedAt = utilities.Times()
	queryInsert := `INSERT INTO users (name, email, password, token, gender, address, phone, role, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
	sqlError := config.DbConn.Sql.QueryRow(queryInsert, request.Name, request.Email, request.Password, request.Token, request.Gender, request.Address,request.Phone, request.Role, request.CreatedAt).Scan(&request.ID)
	if sqlError != nil {
		return nil, sqlError
	}

	return &request, nil
}
