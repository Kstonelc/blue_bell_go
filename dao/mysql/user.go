package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"web_app/models"
)

func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username =?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}
func InsertUser(user *models.User) (err error) {
	//对密码进行加密
	user.Password = encryptPassword(user.Password)
	//执行SQL语句入库
	sqlStr := `insert into user(user_id, username, password) value(?,?,?)`
	fmt.Println(user.UserID, user.Username, user.Password)
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte("lc.com"))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id , username, password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return errors.New("用户不存在")
	}
	if err != nil {
		//查询数据失败
		return err
	}
	//判断密码
	password := encryptPassword(oPassword)
	if password != user.Password {
		return errors.New("密码错误")
	}
	return
}
