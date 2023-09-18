package response

import (
	"myproject/models/base"
	departmentDB "myproject/models/department/database"
	"myproject/models/employee_detail/database"
	positionDB "myproject/models/position/database"
)

type EmployeeDetailResponse struct {
	base.Model
	CurrentSalary float32                 `json:"current_salary"`
	Department    departmentDB.Department `json:"department"`
	Position      positionDB.Position     `json:"position"`
}

func (employeeDetailResponse *EmployeeDetailResponse) MapEmployeeResponse(employeeDetail database.EmployeeDetail) {
	employeeDetailResponse.Model = employeeDetail.Model
	employeeDetailResponse.Department = employeeDetail.Department
	employeeDetailResponse.Position = employeeDetail.Position
}
