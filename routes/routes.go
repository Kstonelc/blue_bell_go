package routes

import (
	"net/http"
	"web_app/controllers"
	"web_app/logger"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	//注册业务
	r.POST("/signup", controllers.SignUpHandler)
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello ok")
	})
	return r
}
