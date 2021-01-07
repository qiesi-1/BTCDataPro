package models

import (
	"BTCDataPro/db_mysql"
	"fmt"
)

type EmailUser struct {
	Id   int    `form:"id"`
	EmailName string `form:"email_name"`
}

//保存用于邮箱账号
func (e EmailUser) AddEmailUser() (int64,error){
	rs,err := db_mysql.Db.Exec("insert into email_login(email_name)" +
		"values(?)",e.EmailName)
	if err != nil {
		//保存数据出错
		fmt.Println(err.Error())
		return -1,err
	}

	_,err = rs.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		return -1,err
	}
	return -1,err
}