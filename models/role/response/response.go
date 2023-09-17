package response

import (
	"myproject/models/base"
	"myproject/models/role/database"
)

type RoleResponse struct {
	base.Model
	database.Role
}

func (roleResponse *RoleResponse) MapRoleResponse(role database.Role) {
	roleResponse.Model = role.Model
	roleResponse.Name = role.Name
}
