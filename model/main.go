package model

import (
	"sake_io_auth/db"
	"sake_io_auth/typeFile"

	"github.com/sirupsen/logrus"
)

func CreateToken(userInfo typeFile.LoginType) typeFile.Token {
	db := db.MySQL()
	defer db.Close()

	stmt, err := db.Prepare("SELECT id FROM user WHERE mail=? AND pass=?;")
	if err != nil {
		logrus.Error(err)
	}
	rows, err := stmt.Query(userInfo.Mail, userInfo.Pass)
	if err != nil {
		logrus.Error(err)
	}

	user_id := ""
	for rows.Next() {
		if err := rows.Scan(&user_id); err != nil {
			logrus.Error(err)
		}
	}

	return typeFile.Token{"id_sample", "refresh_sample"}
}
