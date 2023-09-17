package controllers

import (
	"errors"
	"myproject/configs"
	"myproject/models/base"
	userDB "myproject/models/user/database"
	"myproject/models/user/request"
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
	userDatabase.Username = userRegister.Username
	userDatabase.Email = userRegister.Email
	userDatabase.Password = userRegister.Password
	userDatabase.RoleID = userRegister.RoleID
	result := configs.DB.Save(&userDatabase)

	if result.Error != nil {
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
	var editUser request.EditUser
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
	res := configs.DB.Model(&userDatabase).Where("id = ?", id).Updates(&editUser).Preload(clause.Associations).First(&userDatabase)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "User not found",
			Data:    nil,
		})
	} else if res.Error != nil {
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
	).First(&userDatabase)

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
	configs.DB.Preload(clause.Associations).Find(&users)

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

	result := configs.DB.Preload(clause.Associations).First(&user)

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
		Status:  false,
		Message: "Success delete user",
		Data:    userResponse,
	})
}
