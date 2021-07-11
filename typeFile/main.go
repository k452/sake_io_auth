package typeFile

type LoginType struct {
	Mail string `json:"mail"`
	Pass string `json:"pass"`
}

type Token struct {
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserID struct {
	UserID string
}
