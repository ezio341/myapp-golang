package response

import "myproject/models/role/database"

type RoleResponse struct {
	database.Role
}

func (roleResponse *RoleResponse) MapRoleResponse(role database.Role) {
	roleResponse.Name = role.Name
}
