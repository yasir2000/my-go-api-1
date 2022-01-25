package services

func CreateUser(user users.User) (*users.User, *error.RestErr) {
	// Validate
	if err := user.Validate(); err != nil {
		return nil, err
	}
	// Save to database
	user.Save
}
