package route

import "github.com/gin-gonic/gin"

func Bind(server *gin.Engine) {
	groupuser := server.Group("/user")
	{
		groupuser.POST("/create", CreateUser())
		groupuser.GET("/get/:name", GetUser())
		groupuser.POST("/update/:name", Updatedate())
	}
}
