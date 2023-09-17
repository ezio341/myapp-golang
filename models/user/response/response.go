package response

import (
	"myproject/models/base"
	"myproject/models/user/database"
)

type UserResponse struct {
	base.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Admin    bool   `json:"admin"`
}

func (userResponse *UserResponse) MapUserResponse(user database.User) {
	userResponse.Model = user.Model
	userResponse.Username = user.Username
	userResponse.Email = user.Email
}
