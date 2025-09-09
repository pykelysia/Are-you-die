package api

import (
	"are-you-die/database"
	"hash/crc32"

	"github.com/gin-gonic/gin"
)

func CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var userdata struct {
			Email    string `json:"email"`
			Username string `json:"name"`
		}
		err := ctx.ShouldBindJSON(&userdata)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid request body"})
			return
		}

		// 创建新用户
		id := crc32.ChecksumIEEE([]byte(userdata.Username))
		user := &database.User{
			ID:       uint(id),
			Username: userdata.Username,
			Email:    userdata.Email,
		}
		err = database.NewUser().Create(user)
		if err != nil {
			ctx.JSON(500, gin.H{"error": "failed to create user"})
			return
		}
	}
}
