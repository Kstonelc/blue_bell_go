package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) {
	//判断用户存不存在
	mysql.QueryUserByName()
	//生成UID
	snowflake.GenID()
	//用户密码加密
	//保存进数据库
	mysql.InsertUser()
}
