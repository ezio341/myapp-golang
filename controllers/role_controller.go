package controllers

import (
	"myproject/configs"
	"myproject/models/base"
	"myproject/models/role/database"
	"myproject/models/role/request"
	"myproject/models/role/response"
	"net/http"

	"github.com/labstack/echo/v4"
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
		Message: "Success add role",
		Data:    roleArrayResponse,
	})
}
