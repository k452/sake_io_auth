package db

import (
	"database/sql"
	"sake_io_auth/conf"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	bolt "go.etcd.io/bbolt"
)

func BBolt() *bolt.DB {
	database, err := bolt.Open("./data.db", 0666, nil)
	if err != nil {
		logrus.Error(err)
	}

	return database
}

func MySQL() *sql.DB {
	DBInfo := conf.GetDBInfo()
	db, err := sql.Open("mysql", DBInfo.USER+":"+DBInfo.PASSWORD+"@tcp("+DBInfo.HOST+":"+DBInfo.PORT+")/"+DBInfo.DB+"?parseTime=true")

	if err != nil {
		logrus.Error(err)
	}

	return db
}
