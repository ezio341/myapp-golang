package response

import "myproject/models/department/database"

type DepartmentResponse struct {
	database.Department
}

func (departmentResponse *DepartmentResponse) MapDepartmentResponse(department database.Department) {
	departmentResponse.Model = department.Model
	departmentResponse.Name = department.Name
}
