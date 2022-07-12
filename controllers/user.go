package controllers

import (
	"fmt"
	"net/http"
	"web_app/logic"
	"web_app/models"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	//1 参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数有误 直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		//判断err是不是validator。ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			//"msg": "请求参数有误",'
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	//手动判断 对请求参数进行校验 下面使用validator库进行参数校验go语言中进行参数校验的库
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
	//	zap.L().Error("SignUp with invalid param")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数有误",
	//	})
	//	return
	//}
	fmt.Println(p)
	//2 业务处理
	logic.SignUp(p)
	//3 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
