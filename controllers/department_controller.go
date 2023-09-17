package controllers

import (
	"errors"
	"myproject/configs"
	"myproject/models/base"
	"myproject/models/department/database"
	"myproject/models/department/request"
	"myproject/models/department/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AddDepartment(c echo.Context) error {
	var request request.AddDepartment
	c.Bind(&request)

	if !request.IsValid() {
		return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status:  false,
			Message: "Invalid request",
			Data:    nil,
		})
	}

	var department database.Department
	department.MapAddDepartment(request)
	result := configs.DB.Create(&department)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Failed add department",
			Data:    nil,
		})
	}

	var departmentResponse response.DepartmentResponse
	departmentResponse.MapDepartmentResponse(department)
	return c.JSON(http.StatusCreated, base.BaseResponse{
		Status:  true,
		Message: "Success add department",
		Data:    departmentResponse,
	})
}

func GetDepartment(c echo.Context) error {
	var department []database.Department
	result := configs.DB.Find(&department)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Failed get department",
			Data:    nil,
		})
	}

	departmentResponse := make([]response.DepartmentResponse, len(department))
	for i, dmt := range department {
		departmentResponse[i].MapDepartmentResponse(dmt)
	}
	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success get department",
		Data:    departmentResponse,
	})
}

func DeleteDepartment(c echo.Context) error {
	id, errConv := strconv.ParseUint(c.Param("id"), 10, strconv.IntSize)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status:  false,
			Message: "Invalid Request",
			Data:    nil,
		})
	}

	databaseDepartment := database.Department{
		Model: base.Model{
			ID: uint(id),
		},
	}
	result := configs.DB.First(&databaseDepartment).Delete(&databaseDepartment)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, base.BaseResponse{
			Status:  false,
			Message: "department not found",
			Data:    nil,
		})
	} else if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Failed delete department",
			Data:    nil,
		})
	}

	var departmentResponse response.DepartmentResponse
	departmentResponse.MapDepartmentResponse(databaseDepartment)

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success Delete department",
		Data:    departmentResponse,
	})
}
