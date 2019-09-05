package user_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest_api_gin/http/dao/user_dao"
	"rest_api_gin/http/dto/user_dto"
	"rest_api_gin/http/models"
	"strconv"
)


func CreateUser(c *gin.Context)  {
	var dto user_dto.CreateUserDTO
	err := c.BindJSON(&dto)
	user, err := user_dao.CreateUser(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, models.UserMapper{
		user.ID,
		user.Name,
		user.Email,
		user.Phone,
		user.ClientType,
		user.CreatedAt,
		user.UpdatedAt,
	})
	return
}

func GetUserById(c *gin.Context)  {
	id := c.Param("id")
	idParseInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		c.Abort()
		return
	}
	user, err := user_dao.GetUserById(idParseInt)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, models.UserMapper{
		user.ID,
		user.Name,
		user.Email,
		user.Phone,
		user.ClientType,
		user.CreatedAt,
		user.UpdatedAt,
	})
	return

}

func GetUserByNameOrEmail(c *gin.Context)  {
	name := c.Param("name")
	user, err := user_dao.GetUserByNameOrEmail(name)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	c.JSON(http.StatusOK, user)
	return
}

func ListUser(c *gin.Context)  {
	result, err := user_dao.ListUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result)
	return
}

func PaginateUser()  {

}

func UpdateUser(c *gin.Context)  {
	var dto user_dto.UpdateUserDTO
	id := c.Param("id")
	idParseInt, err := strconv.Atoi(id)
	errorBindJson := c.BindJSON(&dto)
	if errorBindJson != nil {
		c.JSON(http.StatusInternalServerError, errorBindJson)
		c.Abort()
		return
	}
	user, err := user_dao.UpdateUser(idParseInt, dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, models.UserMapper{
		user.ID,
		user.Name,
		user.Email,
		user.Phone,
		user.ClientType,
		user.CreatedAt,
		user.UpdatedAt,
	})
	return
}

func DeleteUser(c *gin.Context)  {
	id := c.Param("id")
	idParseInt, _ := strconv.Atoi(id)
	err := user_dao.DeleteUser(idParseInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, nil)
	return
}
