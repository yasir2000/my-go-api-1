package users

import "github.com/yasir2000/my-go-api-1/utils/errors"

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email"
)

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Password)
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
