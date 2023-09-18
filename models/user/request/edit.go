package request

import "myproject/models/employee_detail/database"

type EditUser struct {
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
	RoleID   uint   `json:"role_id"`
	database.EmployeeDetail
}

func (editUser *EditUser) IsValid() bool {
	return editUser.Username != ""
}
