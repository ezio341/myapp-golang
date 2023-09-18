package response

import (
	"myproject/models/base"
	employeeDB "myproject/models/employee_detail/database"
	roleDB "myproject/models/role/database"
	userDB "myproject/models/user/database"
)

type UserResponse struct {
	base.Model
	Username       string                    `json:"username"`
	Email          string                    `json:"email"`
	Role           roleDB.Role               `json:"role"`
	EmployeeDetail employeeDB.EmployeeDetail `json:"employee_detail"`
}

func (userResponse *UserResponse) MapUserResponse(user userDB.User) {
	userResponse.Model = user.Model
	userResponse.Username = user.Username
	userResponse.Email = user.Email
	userResponse.Role = user.Role
	userResponse.EmployeeDetail = user.EmployeeDetail
}
