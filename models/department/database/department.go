package database

import (
	"myproject/models/base"
	"myproject/models/department/request"
)

type Department struct {
	base.Model
	Name string `json:"name"`
}

func (department *Department) MapAddDepartment(departmentReq request.AddDepartment) {
	department.Name = departmentReq.Name
}
