package model

import (
	"sake_io_auth/typeFile"
)

func CreateToken(userInfo *typeFile.LoginType) typeFile.Token {
	return typeFile.Token{"id_sample", "refresh_sample"}
}
