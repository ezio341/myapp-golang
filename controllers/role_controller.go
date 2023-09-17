package controllers

import (
	"errors"
	"myproject/configs"
	"myproject/models/base"
	"myproject/models/role/database"
	"myproject/models/role/request"
	"myproject/models/role/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AddRole(c echo.Context) error {
	var addRole request.AddRole
	c.Bind(&addRole)
	if !addRole.IsValid() {
		return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status:  false,
			Message: "Invalid Request",
			Data:    nil,
		})
	}

	var databaseRole database.Role
	databaseRole.MapAddRole(addRole)

	result := configs.DB.Create(&databaseRole)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Internal server error",
			Data:    nil,
		})
	}

	var roleResponse response.RoleResponse
	roleResponse.MapRoleResponse(databaseRole)

	return c.JSON(http.StatusCreated, base.BaseResponse{
		Status:  true,
		Message: "Success add role",
		Data:    roleResponse,
	})
}

func GetRoles(c echo.Context) error {
	var databaseRole []database.Role
	result := configs.DB.Find(&databaseRole)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Internal server error",
			Data:    nil,
		})
	}

	roleArrayResponse := make([]response.RoleResponse, len(databaseRole))
	for i, role := range databaseRole {
		roleArrayResponse[i].MapRoleResponse(role)
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success get role",
		Data:    roleArrayResponse,
	})
}

func DeleteRole(c echo.Context) error {
	id, errConv := strconv.ParseUint(c.Param("id"), 10, strconv.IntSize)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status:  false,
			Message: "Invalid Request",
			Data:    nil,
		})
	}

	databaseRole := database.Role{
		Model: base.Model{
			ID: uint(id),
		},
	}
	result := configs.DB.First(&databaseRole).Delete(&databaseRole)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, base.BaseResponse{
			Status:  false,
			Message: "Role not found",
			Data:    nil,
		})
	} else if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Failed delete role",
			Data:    nil,
		})
	}

	var roleResponse response.RoleResponse
	roleResponse.MapRoleResponse(databaseRole)

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success Delete role",
		Data:    roleResponse,
	})
}
