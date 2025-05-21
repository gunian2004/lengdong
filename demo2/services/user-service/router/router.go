package router

import (
	"github.com/coldchain-traceability-system/services/user-service/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine) {

	// 用户注册路由
	r.POST("/register", controllers.Register)
	//用户登录路由
	r.POST("/login", controllers.Login)
}
