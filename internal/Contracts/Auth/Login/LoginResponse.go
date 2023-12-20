package LoginContract

type LoginResponseBody struct {
	ID           string `json:"_id"`
	Email        string `json:"email"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginResponseCookie struct {
	Neptune         string
	NeptuneReserved string
}
