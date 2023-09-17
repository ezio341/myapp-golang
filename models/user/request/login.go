package request

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (userLogin *UserLogin) IsValid() bool {
	return (userLogin.Email != "" && userLogin.Password != "")
}
