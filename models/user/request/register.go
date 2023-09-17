package request

type UserRegister struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	RoleID   uint   `json:"role_id"`
}

func (userRegister *UserRegister) IsValid() bool {
	return userRegister.Email != "" && userRegister.Password != "" && userRegister.Username != ""
}
