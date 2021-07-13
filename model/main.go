package model

import (
	"sake_io_auth/db"
	"sake_io_auth/typeFile"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

func Login(userInfo typeFile.LoginType) typeFile.Token {
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

	if user_id != "" {
		return typeFile.Token{CreateToken(user_id), "refresh_sample"}
	} else {
		return typeFile.Token{"id_sample", "refresh_sample"}
	}
}

func CreateToken(user_id string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		logrus.Error(err)
		return err.Error()
	}
	return t
}
