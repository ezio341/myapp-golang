package request

type AddRole struct {
	Name string `json:"name"`
}

func (addRole *AddRole) IsValid() bool {
	return addRole.Name != ""
}
