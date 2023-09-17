package response

import (
	"myproject/models/base"
	roleDB "myproject/models/role/database"
	userDB "myproject/models/user/database"
)

type UserResponse struct {
	base.Model
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Role     roleDB.Role `gorm:"foreignKey:RoleID"`
}

func (userResponse *UserResponse) MapUserResponse(user userDB.User) {
	userResponse.Model = user.Model
	userResponse.Username = user.Username
	userResponse.Email = user.Email
	userResponse.Role = user.Role
}
