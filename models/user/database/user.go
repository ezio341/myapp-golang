package database

import (
	"myproject/models/base"
	"myproject/models/role/database"
	"myproject/models/user/request"
)

type User struct {
	base.Model
	Username string        `json:"username"`
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Admin    bool          `json:"admin" gorm:"type:bool"`
	RoleID   uint          `json:"role_id"`
	Role     database.Role `gorm:"foreignKey:RoleID"`
}

func (user *User) MapLogin(userLogin request.UserLogin) {
	user.Email = userLogin.Email
	user.Password = userLogin.Password
}

func (user *User) MapRegister(userRegister request.UserRegister) {
	user.Username = userRegister.Username
	user.Email = userRegister.Email
	user.Password = userRegister.Password
	user.RoleID = userRegister.RoleID
}

func (user *User) MapEdit(userEdit request.EditUser) {
	user.Username = userEdit.Username
	user.Admin = userEdit.Admin
	user.RoleID = userEdit.RoleID
}
