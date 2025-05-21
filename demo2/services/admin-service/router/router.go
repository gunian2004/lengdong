package router

import (
	"github.com/coldchain-traceability-system/services/admin-service/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine) {
	//查看冷冻品路由
	r.GET("/view_product", controllers.GetProduct)
	//查看用户路由
	r.GET("/view_factory", controllers.GetFactory)
	//查看经销商路由
	r.GET("/view_distributor", controllers.GetDistributor)
	//查看零售商路由
	r.GET("/view_retailer", controllers.GetRetailer)
	//查看监管机构路由
	r.GET("/view_regulator", controllers.GetRegulator)
	//查看消费者路由
	r.GET("/view_consumer", controllers.GetConsumer)
	//添加管理员路由
	r.POST("/add_admin", controllers.Admin)
	//查看所有用户路由
	r.GET("/view_user", controllers.GetUser)
	//查看管理员路由
	r.GET("/view_admin", controllers.GetAdmin)
	//修改角色路由
	r.POST("/update_role", controllers.UpdateRole)
	//添加用户路由
	r.POST("/add_user", controllers.AddUser)
	//修改用户审核状态信息路由
	r.POST("/update_user_pending", controllers.UpdateUserPending)
	//修改产品审核状态信息路由
	r.POST("/update_product_pending", controllers.UpdateProductPending)
	//查询审核记录表
	r.GET("/view_audit_log", controllers.GetAuditLog)
}
