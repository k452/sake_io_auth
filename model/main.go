package model

import (
	"github.com/sirupsen/logrus"
	"sake_io_auth/db"
	"sake_io_auth/token"
	"sake_io_auth/typeFile"
	"strconv"
	"time"
)

func Login(userInfo typeFile.LoginType) typeFile.Token {
	db := db.Connect()
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
		idToken, _ := token.CreateIDToken(user_id)
		refreshToken, _ := token.CreateRefreshToken(user_id)
		stmt2, err := db.Prepare("INSERT INTO token(user_id, id_token, refresh_token, created_at, updated_at) VALUES(?, ?, ?, ?, ?);")
		if err != nil {
			logrus.Error(err)
		}
		defer stmt2.Close()

		user_id_int, _ := strconv.Atoi(user_id)
		_, err = stmt2.Exec(user_id_int, idToken, refreshToken, time.Now(), time.Now())
		if err != nil {
			logrus.Error(err)
			panic(err)
		}

		return typeFile.Token{idToken, refreshToken}
	} else {
		return typeFile.Token{"", ""}
	}
}

func AllToken() []typeFile.AllToken {
	db := db.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM token;")
	if err != nil {
		logrus.Error(err)
	}

	var tokens []typeFile.AllToken
	for rows.Next() {
		token := typeFile.AllToken{}
		if err := rows.Scan(&token.TokenID, &token.UserID, &token.IDToken, &token.RefreshToken, &token.Created_at, &token.Updated_at); err != nil {
			logrus.Error(err)
		}
		tokens = append(tokens, token)
	}
	return tokens
}
