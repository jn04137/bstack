package models

type Team struct {
	Id int `json:"id"`
	TeamName string `json:"teamName"`
	OwnerId int `json:"owner"`
}
