package models

type Team struct {
	Id          int    `json:"id"`
	TeamName    string `json:"teamName"`
	TeamNanoId	string `json:"teamNanoId"`
	OwnerId     int    `json:"owner,omitempty"`
	OwnerNanoId string `json:"ownerNanoId"`
	Details     string `json:"details,omitempty"`
}
