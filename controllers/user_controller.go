package controllers

import (
	"errors"
	"fmt"
	"myproject/configs"
	"myproject/models/base"
	employeeDB "myproject/models/employee_detail/database"
	employeeReq "myproject/models/employee_detail/request"
	userDB "myproject/models/user/database"
	"myproject/models/user/request"
	userReq "myproject/models/user/request"
	"myproject/models/user/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func AddUser(c echo.Context) error {
	var userRegister request.UserRegister
	c.Bind(&userRegister)
	if !userRegister.IsValid() {
		return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status:  false,
			Message: "Invalid Request",
			Data:    nil,
		})
	}

	var userDatabase userDB.User
	var employeeDatabase employeeDB.EmployeeDetail

	employeeDatabase.CurrentSalary = userRegister.CurrentSalary
	employeeDatabase.DepartmentID = userRegister.DepartmentID
	employeeDatabase.PositionID = userRegister.PositionID
	resultEmployee := configs.DB.Save(&employeeDatabase)

	userDatabase.Username = userRegister.Username
	userDatabase.Email = userRegister.Email
	userDatabase.Password = userRegister.Password
	userDatabase.RoleID = userRegister.RoleID
	userDatabase.EmployeeDetailID = employeeDatabase.ID

	result := configs.DB.Save(&userDatabase).Preload(clause.Associations).First(&userDatabase)

	if result.Error != nil || resultEmployee.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Failed store user",
			Data:    nil,
		})
	}

	var userResponse response.UserResponse
	userResponse.MapUserResponse(userDatabase)

	return c.JSON(http.StatusCreated, base.BaseResponse{
		Status:  true,
		Message: "Success store user",
		Data:    userResponse,
	})
}
func EditUser(c echo.Context) error {
	var editUser userReq.EditUser
	id, err := strconv.ParseUint(c.Param("id"), 10, strconv.IntSize)
	c.Bind(&editUser)
	if !editUser.IsValid() || err != nil {
		return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status:  false,
			Message: "Invalid Request",
			Data:    nil,
		})
	}

	var userDatabase userDB.User
	userDatabase.ID = uint(id)
	userDatabase.MapEdit(editUser)
	resUser := configs.DB.Model(&userDatabase).Updates(&userDatabase).Preload(clause.Associations).First(&userDatabase)

	var employeeDatabase employeeDB.EmployeeDetail
	employeeDatabase.ID = userDatabase.EmployeeDetail.ID
	fmt.Print(editUser)
	employeeDatabase.MapEditEmployeeDetail(employeeReq.EditEmployeeDetail{
		DepartmentID:  editUser.DepartmentID,
		PositionID:    editUser.PositionID,
		CurrentSalary: editUser.CurrentSalary,
	})

	resEmployee := configs.DB.Model(&employeeDatabase).Updates(&employeeDatabase).Preload(clause.Associations).First(&employeeDatabase)
	userDatabase.EmployeeDetail = employeeDatabase

	if errors.Is(resUser.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "User not found",
			Data:    nil,
		})
	} else if resUser.Error != nil && resEmployee.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Failed update user",
			Data:    nil,
		})
	}

	var userResponse response.UserResponse
	userResponse.MapUserResponse(userDatabase)

	return c.JSON(http.StatusCreated, base.BaseResponse{
		Status:  true,
		Message: "Success update user",
		Data:    userResponse,
	})
}

func Login(c echo.Context) error {
	var login request.UserLogin
	c.Bind(&login)

	if !login.IsValid() {
		return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status:  false,
			Message: "Invalid Request",
			Data:    nil,
		})
	}

	var userDatabase userDB.User
	userDatabase.MapLogin(login)

	result := configs.DB.Where(
		"email = ? AND password= ?",
		userDatabase.Email,
		userDatabase.Password,
	).Preload(clause.Associations).First(&userDatabase)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusUnauthorized, base.BaseResponse{
			Status:  false,
			Message: "Login failed check email and password",
			Data:    nil,
		})
	}

	var loginResponse response.LoginResponse
	loginResponse.MaploginResponse(userDatabase)

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Login success",
		Data:    loginResponse,
	})
}

func GetUsers(c echo.Context) error {
	var users []userDB.User
	configs.DB.Preload(clause.Associations).Preload("EmployeeDetail.Department").Preload("EmployeeDetail.Position").Find(&users)

	userArrayResponse := make([]response.UserResponse, len(users))
	for i, user := range users {
		userArrayResponse[i].MapUserResponse(user)
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success get users",
		Data:    userArrayResponse,
	})
}

func GetUser(c echo.Context) error {
	id, paramErr := strconv.ParseUint(c.Param("id"), 10, strconv.IntSize)
	if paramErr != nil {
		return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status:  false,
			Message: "Invalid Request",
			Data:    nil,
		})
	}

	var user userDB.User
	user.ID = uint(id)

	result := configs.DB.Preload(clause.Associations).Preload("EmployeeDetail.Department").Preload("EmployeeDetail.Position").First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, base.BaseResponse{
				Status:  false,
				Message: "User not found",
				Data:    nil,
			})
		}
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Failed get user",
			Data:    nil,
		})
	}

	var userResponse response.UserResponse
	userResponse.MapUserResponse(user)

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success get user",
		Data:    user,
	})
}

func DeleteUser(c echo.Context) error {
	id, paramErr := strconv.ParseUint(c.Param("id"), 10, strconv.IntSize)
	if paramErr != nil {
		return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status:  false,
			Message: "Invalid Request",
			Data:    nil,
		})
	}
	user := userDB.User{
		Model: base.Model{
			ID: uint(id),
		},
	}

	result := configs.DB.Preload(clause.Associations).First(&user).Delete(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, base.BaseResponse{
			Status:  false,
			Message: "User not found",
			Data:    nil,
		})
	} else if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Failed delete user",
			Data:    nil,
		})
	}

	var userResponse response.UserResponse
	userResponse.MapUserResponse(user)

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success delete user",
		Data:    userResponse,
	})
}
