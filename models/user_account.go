package models

type UserAccount struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Email string `json:"email,omitempty"`
	NanoId string `json:"nanoId"`
}
