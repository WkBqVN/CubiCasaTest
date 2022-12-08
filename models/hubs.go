package models

type Hub struct {
	HubID     uint   `json:"hubId" gorm:"primaryKey;OnDelete:Cascade"`
	HubName   string `json:"hubName"`
	HubRegion string `json:"hubRegion"`
}
