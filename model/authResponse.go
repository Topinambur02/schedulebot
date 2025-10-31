package model

type AuthResponse struct {
	Token      string `json:"token"`
	Guid       string `json:"guid"`
	JWT        string `json:"jwt"`
	JWTRefresh string `json:"jwt_refresh"`
}