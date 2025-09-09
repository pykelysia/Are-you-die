package api

import (
	"are-you-die/database"
	"hash/crc32"

	"github.com/gin-gonic/gin"
)

func Updatedate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从请求参数中获取用户ID
		name := ctx.Param("name")
		if name == "" {
			ctx.JSON(400, gin.H{"error": "name is required"})
			return
		}

		id := crc32.ChecksumIEEE([]byte(name))

		// 查询用户信息
		_, err := database.NewUser().Get(uint(id))
		if err != nil {
			ctx.JSON(500, gin.H{"error": "failed to get user"})
			return
		}

		// 更新用户的最新一次签到时间
		err = database.NewDate().UpdateCheckin(uint(id))
		if err != nil {
			ctx.JSON(500, gin.H{"error": "failed to update checkin date"})
			return
		}
		ctx.JSON(200, gin.H{"msg": "checkin updated"})
	}
}
