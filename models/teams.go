package models

type Team struct {
	TeamID   uint   `json:"teamId" gorm:"primaryKey"`
	TeamName string `json:"teamName"`
	HubID    uint   `json:"hubId" gorm:"foreignKey:hub_id;constraint:OnDelete:Cascade;references:hub_id"`
}
