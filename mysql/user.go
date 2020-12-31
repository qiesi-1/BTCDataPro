package mysql

import (
	"data/db_mysql"
	"fmt"
)

type Info struct {
	Hash string `form:"hash"`
	Height int `form:"height"`
	Difficult int `form:difficult`
	Count int `form:count`
	Chain string `form:"chain"`
}

func (i Info) AddUsers() (int64, error) {
	rs, err := db_mysql.Db.Exec("insert into register(height, difficult, count, chain)" +
		"value(?,?,?,?)", i.Height, i.Difficult, i.Count, i.Chain)
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

func (i *Info) QueryUsers() (*Info, error) {
	row := db_mysql.Db.QueryRow("select hash, height, difficult, count, chain from register where hash = ? and height = ?",
	    i.Hash, i.Height)

	err := row.Scan(&i.Hash, &i.Height, &i.Difficult, &i.Count, &i.Chain)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return i, err
}

