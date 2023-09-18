package database

import (
	"myproject/models/base"
	"myproject/models/position/request"
)

type Position struct {
	base.Model
	Name       string  `json:"name"`
	BaseSalary float32 `json:"base_salary"`
}

func (position *Position) MapAddPosition(addPosition request.AddPosition) {
	position.Name = addPosition.Name
	position.BaseSalary = addPosition.BaseSalary
}
