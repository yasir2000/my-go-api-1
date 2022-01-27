package services

import (
	"github.com/yasir2000/my-go-api-1/domain/users"
	"github.com/yasir2000/my-go-api-1/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	// Validate
	if err := user.Validate(); err != nil {
		return nil, err
	}
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
