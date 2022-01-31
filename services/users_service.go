package services

import (
	"github.com/yasir2000/my-go-api-1/domain/users"
	"github.com/yasir2000/my-go-api-1/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	// Validate
	if err := user.Validate(); err != nil {
		return nil, err
	}
	//encrypt password
	pwSlice, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, errors.NewBadRequestError("failed to encypt the password")
	}
	user.Password = string(pwSlice[:])
	// Save to database
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	current.FirstName = user.FirstName
	current.LastName = user.LastName
	current.Email = user.Email

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func SearchUser(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindbyStatus(status)
}

func GetbyEmail(user users.User) (*users.User, *errors.RestErr) {
	result := &users.User{Email: user.Email}
	if err := result.GetByEmail(); err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
		return nil, errors.NewBadRequestError("failed to decrypt password")
	}

	resultWp := &users.User{Id: result.Id, FirstName: result.FirstName, LastName: result.LastName, Email: result.Email}
	return resultWp, nil
}
