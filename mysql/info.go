package mysql

import (
	"data/db_mysql"
	"fmt"
)

type Info struct {
	Hash string `form:"hash"`
	Height int `form:"height"`
	Difficult int `form:difficult`
	Size int `from:size`
	Count int `form:count`
	Chain string `form:"chain"`
}

func (i Info) AddInfo() (int64, error) {
	rs, err := db_mysql.Db.Exec("insert into rpc_info(hash, height, difficult, size, count, chain)" +
		"value(?,?,?,?)", i.Hash, i.Height, i.Difficult, i.Size, i.Count, i.Chain)
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

func (i *Info) QueryInfo() (*Info, error) {
	row := db_mysql.Db.QueryRow("select hash, height, difficult, size, count, chain from rpc_info where hash = ? and height = ?",
	    i.Hash, i.Height)

	err := row.Scan(&i.Hash, &i.Height, &i.Difficult, &i.Size, &i.Count, &i.Chain)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return i, err
}

