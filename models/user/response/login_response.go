package response

import (
	"myproject/middlewares"
	"myproject/models/base"
	roleDB "myproject/models/role/database"
	"myproject/models/user/database"
)

type LoginResponse struct {
	base.Model
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Token    string      `json:"token"`
	Role     roleDB.Role `json:"role"`
}

func (loginResponse *LoginResponse) MaploginResponse(user database.User) {
	loginResponse.Model = user.Model
	loginResponse.Username = user.Username
	loginResponse.Email = user.Email
	loginResponse.Role = user.Role
	loginResponse.Token = middlewares.GenerateToken(user)
}
