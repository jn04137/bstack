package models

type Team struct {
	Id          int    `json:"id"`
	TeamName    string `json:"teamName"`
	TeamNanoId	string `json:"teamNanoId"`
	OwnerId     int    `json:"owner,omitempty"`
	OwnerNanoId string `json:"ownerNanoId,omitempty"`
	Details     string `json:"teamDetails,omitempty"`
	OwnerName 	string `json:"ownerUsername,omitempty"`
}
