package user_dao

import (
	"github.com/gin-gonic/gin"
	"rest-api-gin/configs"
	"rest-api-gin/http/dto/user_dto"
	"rest-api-gin/http/models"
	"rest-api-gin/utils"
)

func CreateUser(dto user_dto.CreateUserDTO) (user models.User, error error) {
	db := configs.Database()
	password, _ := utils.HashPassword(dto.Password)
	user = models.User{
		Name:       dto.Name,
		Email:      dto.Email,
		Phone:      dto.Phone,
		Password:   password,
		ClientType: dto.ClientType,
	}
	err := db.Create(&user).Error
	if err != nil {
		error = err
	}
	return user, error
}

func GetUserById(id int) (user models.User, error error) {
	db := configs.Database()
	err := db.Find(&user, id).Error
	if err != nil {
		error = err
	}
	return user, error
}

func GetUserByNameOrEmail(name string) (user models.User, error error)  {
	db := configs.Database()
	err := db.Find(&user, name).Error
	if err != nil {
		error = err
	}
	return user, error
}

func ListUser() (result gin.H, error error) {
	var user []models.User
	db := configs.Database()
	db.Find(&user)
	result = gin.H{"data": user}
	return result, error
}

func UpdateUser(id int, dto user_dto.UpdateUserDTO) (user models.User, error error) {
	db := configs.Database()
	password, _ := utils.HashPassword(dto.Password)
	err := db.First(&user, id).Error
	if err != nil {
		error = err
	}
	user.Name = dto.Name
	user.Email = dto.Email
	user.Password = password
	user.ClientType = dto.ClientType
	errorSaveToDb := db.Save(&user).Error
	if errorSaveToDb != nil {
		error = errorSaveToDb
	}
	return user, error
}

func DeleteUser(id int) (error error) {
	var user models.User
	db := configs.Database()
	errorGetId := db.First(&user, id).Delete(&user).Error
	if errorGetId != nil {
		error = errorGetId
	}
	return error
}