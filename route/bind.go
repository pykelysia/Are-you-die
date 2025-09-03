package route

import "github.com/gin-gonic/gin"

func Bind(server *gin.Engine) {
	server.GET("/getuser", GetUser())
	server.POST("/updatedate/:name", Updatedate())
}
