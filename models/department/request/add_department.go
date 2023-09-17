package request

type AddDepartment struct {
	Name string `json:"name"`
}

func (addDepartment *AddDepartment) IsValid() bool {
	return addDepartment.Name != ""
}
