package database

import (
	"myproject/models/base"
	"myproject/models/role/request"
)

type Role struct {
	base.Model
	Name string `json:"name"`
}

func (role *Role) MapAddRole(addRole request.AddRole) {
	role.Name = addRole.Name
}
