package database

import (
	"myproject/models/base"
	employeeDB "myproject/models/employee_detail/database"
	employeeReq "myproject/models/employee_detail/request"
	roleDB "myproject/models/role/database"
	userReq "myproject/models/user/request"
)

type User struct {
	base.Model
	Username         string                    `json:"username"`
	Email            string                    `json:"email" gorm:"index:,unique"`
	Password         string                    `json:"password"`
	RoleID           uint                      `json:"role_id"`
	Role             roleDB.Role               `gorm:"foreignKey:RoleID"`
	EmployeeDetailID uint                      `json:"employee_detail_id"`
	EmployeeDetail   employeeDB.EmployeeDetail `gorm:"foreignKey:EmployeeDetailID"`
}

func (user *User) MapLogin(userLogin userReq.UserLogin) {
	user.Email = userLogin.Email
	user.Password = userLogin.Password
}

func (user *User) MapRegister(userRegister userReq.UserRegister) {
	user.Username = userRegister.Username
	user.Email = userRegister.Email
	user.Password = userRegister.Password
	user.RoleID = userRegister.RoleID
	var employee employeeDB.EmployeeDetail
	employee.MapAddEmployeeDetail(employeeReq.AddEmployeeDetail{
		CurrentSalary: userRegister.CurrentSalary,
		DepartmentID:  userRegister.DepartmentID,
		PositionID:    userRegister.PositionID,
	})
}

func (user *User) MapEdit(userEdit userReq.EditUser) {
	user.Username = userEdit.Username
	user.RoleID = userEdit.RoleID
	var employee employeeDB.EmployeeDetail
	employee.MapAddEmployeeDetail(employeeReq.AddEmployeeDetail{
		CurrentSalary: userEdit.CurrentSalary,
		DepartmentID:  userEdit.DepartmentID,
		PositionID:    userEdit.PositionID,
	})
}
