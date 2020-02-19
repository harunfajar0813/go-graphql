package datastore

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewMyqlDB() (*gorm.DB, error) {
	DBMS := "mysql"
	mySqlConfig := &mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1",
		DBName:               "cobain_dulu",
		AllowNativePasswords: true,
		Params: map[string]string{
			"parseTime": "true",
		},
	}

	return gorm.Open(DBMS, mySqlConfig.FormatDSN())
}
