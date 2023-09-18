package database

import (
	"myproject/models/base"
	departmentDB "myproject/models/department/database"
	"myproject/models/employee_detail/request"
	positionDB "myproject/models/position/database"
)

type EmployeeDetail struct {
	base.Model
	DepartmentID  uint                    `json:"department_id"`
	PositionID    uint                    `json:"position_id"`
	CurrentSalary float32                 `json:"current_salary"`
	Department    departmentDB.Department `json:"department" gorm:"foreignKey:DepartmentID"`
	Position      positionDB.Position     `json:"position" gorm:"foreignKey:PositionID"`
}

func (employeeDetail *EmployeeDetail) MapAddEmployeeDetail(request request.AddEmployeeDetail) {
	employeeDetail.DepartmentID = request.DepartmentID
	employeeDetail.PositionID = request.PositionID
	employeeDetail.CurrentSalary = request.CurrentSalary
}
func (employeeDetail *EmployeeDetail) MapEditEmployeeDetail(request request.EditEmployeeDetail) {
	employeeDetail.DepartmentID = request.DepartmentID
	employeeDetail.PositionID = request.PositionID
	employeeDetail.CurrentSalary = request.CurrentSalary
}
