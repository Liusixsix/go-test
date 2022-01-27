package routes

import (
	"gin-demo/common"
	"gin-demo/controller"
	"gin-demo/middleware"
	"gin-demo/model"
	"gin-demo/utils"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.POST("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	categoryRoutes := r.Group("/category")
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("", categoryController.Create)
	categoryRoutes.PUT("/:id", categoryController.Update)
	categoryRoutes.GET("/:id", categoryController.Show)
	categoryRoutes.DELETE("/:id", categoryController.Delete)
	InitPachong()

	r.GET("pachong", utils.Reptile)
	r.GET("/pachong/list", controller.GetList)
	return r
}

func InitPachong() {
	db := common.GetDb()
	db.AutoMigrate(model.Pachong{})
}
