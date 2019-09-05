package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"rest_api_gin/configs"
	"rest_api_gin/http/controller/user_controller"
)

func main()  {
	configs.DotEnv()

	router := gin.Default()
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			//auth := v1.Group("/auth")
			//{
				//auth.POST("/register", auth_controller.Register)
				//auth.POST("/login", auth_controller.Login)
			//}
			user := v1.Group("/users")
			{
				user.POST("", user_controller.CreateUser)
				user.GET("", user_controller.ListUser)
				user.GET("/:id", user_controller.GetUserById)
				user.PUT("/:id", user_controller.UpdateUser)
				user.DELETE("/:id", user_controller.DeleteUser)
			}

		}
	}
	_ = router.Run(os.Getenv("APP_PORT"))

}
