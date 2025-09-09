package route

import (
	"are-you-die/api"

	"github.com/gin-gonic/gin"
)

func Bind(server *gin.Engine) {
	groupuser := server.Group("/user")
	{
		groupuser.POST("/create", api.CreateUser())
		groupuser.GET("/get/:name", api.GetUser())
		groupuser.POST("/update/:name", api.Updatedate())
	}
}
