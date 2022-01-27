package users

import (
	"github.com/yasir2000/my-go-api-1/datasource/mysql/users_db"
	"github.com/yasir2000/my-go-api-1/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users_db_01.users(first_name, last_name, email, password) values(?,?,?,?);"
	queryGetUser    = "SELECT Id, first_name, last_name, email FROM users WHERE id=?;"
)

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password)
	if saveErr != nil {
		return errors.NewInternalServerError("database error")
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	user.Id = userId
	return nil
}

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email); getErr != nil {
		return errors.NewInternalServerError("database error")
	}
	return nil
}
