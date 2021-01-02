package models

import (
	"BTCDataPro/db_mysql"
	"fmt"
)

type Info struct {
	Hash      string `json:"hash"`
	Height    int    `json:"height"`
	Difficult int    `json:"difficult"`
	Size      int    `json:"size"`
	Count     int    `json:"count"`
	Chain     string `json:"chain"`
}

func (i Info) AddRpcInfo() (int64, error) {
	rs, err := db_mysql.Db.Exec("insert into rpc_info(hash,height,difficult,size,count,chain)" +
		"value(?,?,?,?,?,?)", i.Hash,i.Height, i.Difficult,i.Size, i.Count, i.Chain)
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

//func (i *Info) QueryRpcInfo() (*Info, error) {
//	row := db_mysql.Db.QueryRow("select hash, height, difficult, count, chain from register where hash = ? and height = ?",
//	    i.Hash, i.Height)
//
//	err := row.Scan(&i.Hash, &i.Height, &i.Difficult, &i.Count, &i.Chain)
//	if err != nil {
//		fmt.Println(err.Error())
//		return nil, err
//	}
//	return i, err
//}

