package request

type AddEmployeeDetail struct {
	CurrentSalary float32 `json:"current_salary"`
	DepartmentID  uint    `json:"department_id"`
	PositionID    uint    `json:"position_id"`
}
