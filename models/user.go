package models

import (
	"BTCDataPro/db_mysql"
	"fmt"
)

type Users struct {
	Id       int    `form:"id"`
	UserName string `form:"username"`
	Pwd      string `form:"pwd"`
	Rpwd     string `form:"rpwd"`
	Phone    string `form:"phone"`
}

//保存用户注册数据
func (u Users) AddUsers() (int64,error){
	rs,err := db_mysql.Db.Exec("insert into register(username,pwd,rpwd,phone)" +
		"values(?,?,?,?)",u.UserName,u.Pwd,u.Rpwd,u.Phone)
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

//查询用户信息
func (u *Users) QueryUsers() (*Users, error){
	row := db_mysql.Db.QueryRow("select id,username,pwd,rpwd,phone from register where username = ? and pwd = ?",
		u.UserName,u.Pwd)

	err := row.Scan(&u.Id,&u.UserName,&u.Pwd,&u.Rpwd,&u.Phone)
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}
	return u,err
}

//根据手机号查询用户有没有注册
func (u Users) QueryUserByphone() (*Users, error) {
	row := db_mysql.Db.QueryRow("select id,username,pwd,phone from register where phone = ? ", u.Phone)
	err := row.Scan(&u.Id,&u.UserName,&u.Pwd,&u.Phone)
	if err != nil {
		return nil, err
	}
	return &u, err
}