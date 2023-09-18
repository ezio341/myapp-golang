package controllers

import (
	"errors"
	"myproject/configs"
	"myproject/models/base"
	"myproject/models/position/database"
	"myproject/models/position/request"
	"myproject/models/position/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AddPosition(c echo.Context) error {
	var addPosition request.AddPosition
	c.Bind(&addPosition)
	if !addPosition.IsValid() {
		return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status:  false,
			Message: "Invalid Request",
			Data:    nil,
		})
	}

	var databasePosition database.Position
	databasePosition.MapAddPosition(addPosition)

	result := configs.DB.Create(&databasePosition)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Internal server error",
			Data:    nil,
		})
	}

	var PositionResponse response.PositionResponse
	PositionResponse.MapPositionResponse(databasePosition)

	return c.JSON(http.StatusCreated, base.BaseResponse{
		Status:  true,
		Message: "Success add position",
		Data:    PositionResponse,
	})
}
func GetPositions(c echo.Context) error {
	var databasePosition []database.Position
	result := configs.DB.Find(&databasePosition)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Internal server error",
			Data:    nil,
		})
	}

	PositionArrayResponse := make([]response.PositionResponse, len(databasePosition))
	for i, position := range databasePosition {
		PositionArrayResponse[i].MapPositionResponse(position)
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success get position",
		Data:    PositionArrayResponse,
	})
}
func DeletePosition(c echo.Context) error {
	id, errConv := strconv.ParseUint(c.Param("id"), 10, strconv.IntSize)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status:  false,
			Message: "Invalid Request",
			Data:    nil,
		})
	}

	databasePosition := database.Position{
		Model: base.Model{
			ID: uint(id),
		},
	}
	result := configs.DB.First(&databasePosition).Delete(&databasePosition)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, base.BaseResponse{
			Status:  false,
			Message: "Position not found",
			Data:    nil,
		})
	} else if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Failed delete position",
			Data:    nil,
		})
	}

	var PositionResponse response.PositionResponse
	PositionResponse.MapPositionResponse(databasePosition)

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success Delete position",
		Data:    PositionResponse,
	})
}
