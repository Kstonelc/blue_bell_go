package routes

import (
	"net/http"
	"web_app/controllers"
	"web_app/logger"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)gin设置成发布模式 就不会在控制台输出日志了
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	//注册业务
	r.POST("/signup", controllers.SignUpHandler)
	r.POST("/login", controllers.LoginHandler)
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello ok")
	})
	return r
}
