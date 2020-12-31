package mysql

import (
	"data/db_mysql"
	"fmt"
)

type Users struct {
	Id int `form:"id"`
	UserName string `form:"username"`
	Pwd string `form:pwd`
	Rpwd string `form:rpwd`
	Phone string `form:"phone"`
}

func (u Users) AddUsers() (int64, error) {
	rs, err := db_mysql.Db.Exec("insert into register(username, pwd, rpwd, phone)" +
		"value(?,?,?,?)", u.UserName, u.Pwd, u.Rpwd, u.Phone)
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}

	_, err = rs.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}
	return -1, err
}

func (u *Users) QueryUsers() (*Users, error) {
	row := db_mysql.Db.QueryRow("select id, username, pwd, phone grom register where username = ? and pwd = ?",
	    u.UserName, u.Pwd)

	err := row.Scan(&u.Id, &u.UserName, &u.Pwd, &u.Rpwd, &u.Phone)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return u, err
}

func (u Users) QueryUserByPhone() (*Users, error) {
	row := db_mysql.Db.QueryRow("select id,username,pwd,phone from register where phone = ? ", u.Phone)
	err := row.Scan(&u.Id, &u.UserName, &u.Pwd, &u.Phone)
	if err != nil {
		return nil, err
	}
	return &u, err
}
