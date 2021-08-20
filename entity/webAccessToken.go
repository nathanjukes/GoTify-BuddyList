package entity

type WebAccessToken struct {
	ClientId                         string `json:"clientId"`
	AccessToken                      string `json:"accessToken"`
	AccessTokenExpirationTimestampMs string `json:"accessTokenExpirationTimestampMs"`
	IsAnonymous                      string `json:"isAnonymous"`
}
