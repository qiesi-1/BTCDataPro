package db_mysql

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

//连接数据库
func DbConnect() {
	fmt.Println("连接mysql数据库成功")
	config := beego.AppConfig //定义config变量,接收并赋值为全局配置变量

	dbDriver := config.String("db_driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbPort := config.String("db_port")
	dbName := config.String("db_name")

	//连接数据库
	connUrl := dbUser + ":" + dbPassword + "@tcp(" + dbIp + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	db, err := sql.Open(dbDriver, connUrl)
	if err != nil { // err不为nil，表示连接数据库时出现了错误, 程序就在此中断就可以，不用再执行了。
		//早解决，早解决
		fmt.Println(err.Error())
		panic("数据库连接错误，请检查配置")
	}
	Db = db
}
