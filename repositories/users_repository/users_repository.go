package usersrepository

import (
	"online-store/config"
	"online-store/models"
	"online-store/utilities"
)

func CreateUser(request models.User) (*models.User, error) {
	request.CreatedAt = utilities.Times()
	queryInsert := `INSERT INTO users (name, email, password, token, gender, address, phone, role, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
	sqlError := config.DbConn.Sql.QueryRow(queryInsert, request.Name, request.Email, request.Password, request.Token, request.Gender, request.Address, request.Phone, request.Role, request.CreatedAt).Scan(&request.ID)
	if sqlError != nil {
		return nil, sqlError
	}

	return &request, nil
}
func ChangeIsLogin(userId, token string, isLogin bool) (int64, error) {
	queryUpdate := `UPDATE users SET is_login = $1, token = $2 WHERE id = $3`
	result, sqlError := config.DbConn.Sql.Exec(queryUpdate, isLogin, token, userId)
	if sqlError != nil {
		return 0, sqlError
	}

	totalRowsUpdated, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return totalRowsUpdated, nil
}

func GetUser(userId, Email, token string) (*models.User, error) {
	var user models.User
	var queryWhere string
	if userId != "" {
		queryWhere = `AND id = '` + userId + `'`
	} else if Email != "" {
		queryWhere = `AND email = '` + Email + `'`
	} else if token != "" {
		queryWhere = `AND token = '` + token + `'`
	}
	querySelect := `SELECT id, name, email, password, token, is_login, gender, address, phone, role, created_at, updated_at, deleted_at FROM users WHERE deleted_at IS NULL ` + queryWhere
	sqlError := config.DbConn.Sql.QueryRow(querySelect).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Token, &user.IsLogin, &user.Gender, &user.Address, &user.Phone, &user.Role, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if sqlError != nil {
		return nil, sqlError
	}

	return &user, nil
}
