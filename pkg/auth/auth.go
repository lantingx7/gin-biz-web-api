package auth

import (
	"errors"

	"github.com/gin-gonic/gin"

	"gin-biz-web-api/model/user_model"
	"gin-biz-web-api/pkg/logger"
)

// CurrentUser 从 gin.context 中获取当前登录的用户信息
func CurrentUser(c *gin.Context) user_model.User {
	user, ok := c.MustGet("current_user_info").(user_model.User)
	if !ok {
		logger.LogErrorIf(errors.New("无法获取当前用户信息"))
		return user_model.User{}
	}

	return user
}

// CurrentUserID 从 gin.context 中获取当前登录的用户 ID
func CurrentUserID(c *gin.Context) uint {
	return c.GetUint("current_user_id")
}
