package typeFile

import (
	"time"
)

type LoginType struct {
	Mail string `json:"mail"`
	Pass string `json:"pass"`
}

type Token struct {
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserID struct {
	UserID string
}

type AllToken struct {
	TokenID      int       `json:"token_id"`
	UserID       int       `json:"user_id"`
	IDToken      string    `json:"id_token"`
	RefreshToken string    `json:"refresh_token"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
}
