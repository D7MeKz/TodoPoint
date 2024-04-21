package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"todopoint/common/auth"
	"todopoint/common/netutils/middleware"
	_ "todopoint/member/docs"
	"todopoint/member/router/controller"
)

// NewMemberRouter
// @title Member API
// @Version 1.0
// @host localhost:3000
func NewMemberRouter(controller *controller.MemberController) *gin.Engine {
	engine := gin.Default()
	// engine.Use(middleware.SetHeader())
	engine.Use(middleware.ErrorHandler())
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router := engine.Group("/members")
	{
		router.Use(auth.TokenAuthMiddleware())
		router.GET("/:memId/valid", controller.IsValidMember)

	}

	authRouter := engine.Group("/auth")
	{
		authRouter.POST("/login", controller.LoginMember)
		authRouter.POST("/register", controller.RegisterMember)
		authRouter.POST("/token", controller.RefreshToken)
	}

	return engine
}
