package response

import (
	"myproject/models/base"
	"myproject/models/position/database"
)

type PositionResponse struct {
	base.Model
	Name       string  `json:"name"`
	BaseSalary float32 `json:"base_salary"`
}

func (positionResponse *PositionResponse) MapPositionResponse(position database.Position) {
	positionResponse.Model = position.Model
	positionResponse.Name = position.Name
	positionResponse.BaseSalary = position.BaseSalary
}
