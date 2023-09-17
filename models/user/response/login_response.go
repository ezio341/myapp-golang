package response

import (
	"myproject/middlewares"
	"myproject/models/base"
	"myproject/models/user/database"
)

type LoginResponse struct {
	base.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Admin    bool   `json:"admin"`
	Token    string `json:"token"`
}

func (loginResponse *LoginResponse) MaploginResponse(user database.User) {
	loginResponse.Model = user.Model
	loginResponse.Username = user.Username
	loginResponse.Email = user.Email
	loginResponse.Admin = user.Admin
	loginResponse.Token = middlewares.GenerateToken(user)
}
