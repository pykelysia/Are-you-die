package route

import (
	"are-you-die/database"
	"hash/crc32"

	"github.com/gin-gonic/gin"
)

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从请求参数中获取用户ID
		name := ctx.Param("name")
		if name == "" {
			ctx.JSON(400, gin.H{"error": "name is required"})
			return
		}

		// 获取id
		id := crc32.ChecksumIEEE([]byte(name))

		// 查询用户信息
		user, err := database.NewUser().Get(uint(id))
		if err != nil {
			ctx.JSON(500, gin.H{"error": "failed to get user"})
			return
		}
		ctx.JSON(200, gin.H{"user": user})
	}
}
