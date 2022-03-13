package user

type UserRegisteredFormatter struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UserLoggedFormatter struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

func FormatUserRegistered(user User) UserRegisteredFormatter {
	userFormatted := UserRegisteredFormatter{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Role: user.Role,
	}

	return userFormatted
}

func FormatUserLogged(user User, tokenGenerated string) UserLoggedFormatter {
	userFormatted := UserLoggedFormatter{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Role: user.Role,
		Token: tokenGenerated,
	}

	return userFormatted
}