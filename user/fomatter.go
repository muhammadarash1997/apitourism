package user

type userRegisteredFormatter struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type userLoggedFormatter struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

func FormatUserRegistered(user User) userRegisteredFormatter {
	userFormatted := userRegisteredFormatter{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Role: user.Role,
	}

	return userFormatted
}

func FormatUserLogged(user User, tokenGenerated string) userLoggedFormatter {
	userFormatted := userLoggedFormatter{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Role: user.Role,
		Token: tokenGenerated,
	}

	return userFormatted
}