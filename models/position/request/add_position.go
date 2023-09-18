package request

type AddPosition struct {
	Name       string  `json:"name"`
	BaseSalary float32 `json:"base_salary" gorm:"default:0"`
}

func (addPosition *AddPosition) IsValid() bool {
	return addPosition.Name != ""
}
