package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/dbalpha/go-ldap-admin/controller"
	"github.com/dbalpha/go-ldap-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDeployRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	deploy := r.Group("/deploy")
	deploy.Use(authMiddleware.MiddlewareFunc())
	deploy.Use(middleware.CasbinMiddleware())
	{
		deploy.GET("/info", controller.Deploy.Info)
		deploy.GET("/list", controller.Deploy.List)
		deploy.POST("/add", controller.Deploy.Add)
		deploy.POST("/delete", controller.Deploy.Delete)
		deploy.POST("/update", controller.Deploy.Update)
		deploy.POST("/history/delete", controller.Deploy.DeployHisDelete)
		deploy.POST("/history/list", controller.Deploy.DeployHisList)
		deploy.POST("/deploy/history/info ", controller.Deploy.DeployHisInfo)

	}
	return r
}
