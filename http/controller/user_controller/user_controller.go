package user_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api-gin/http/dao/user_dao"
	"rest-api-gin/http/dto/user_dto"
	"rest-api-gin/http/response"
	"strconv"
)


func CreateUser(c *gin.Context)  {
	var dto user_dto.CreateUserDTO
	httpMethod := http.MethodPost
	err := c.BindJSON(&dto)
	user, err := user_dao.CreateUser(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.InternalServerErrorFullResponse(httpMethod, err.Error()))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, response.SingleFullResponse(httpMethod, user))
	return
}

func GetUserById(c *gin.Context)  {
	id := c.Param("id")
	httpMethod := http.MethodGet
	idParseInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.InternalServerErrorFullResponse(httpMethod, err.Error()))
		c.Abort()
		return
	}
	user, err := user_dao.GetUserById(idParseInt)
	if err != nil {
		c.JSON(http.StatusNotFound, response.CustomFullResponse(
			http.StatusNotFound,
			httpMethod,
			"Not Found",
			"0"))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, response.SingleFullResponse(httpMethod, user))
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
	httpMethod := http.MethodGet
	result, err := user_dao.ListUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, response.ListFullResponse(httpMethod, result))
	return
}

func PaginateUser()  {
	//TODO
}

func UpdateUser(c *gin.Context)  {
	var dto user_dto.UpdateUserDTO
	httpMethod := http.MethodPut
	id := c.Param("id")
	idParseInt, err := strconv.Atoi(id)
	errorBindJson := c.BindJSON(&dto)
	if errorBindJson != nil {
		c.JSON(http.StatusInternalServerError, response.ValidationFullResponse(httpMethod, errorBindJson.Error()))
		c.Abort()
		return
	}
	user, err := user_dao.UpdateUser(idParseInt, dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, response.Single(user))
	return
}

func DeleteUser(c *gin.Context)  {
	id := c.Param("id")
	httpMethod := http.MethodDelete
	idParseInt, _ := strconv.Atoi(id)
	err := user_dao.DeleteUser(idParseInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.InternalServerErrorFullResponse(httpMethod, err.Error()))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, response.SuccessDeleted())
	return
}
