package models

type User struct {
	UserID   uint   `json:"userId" gorm:"primaryKey"`
	UserName string `json:"userName"`
	Role     string `json:"role"`
	TeamId   uint   `json:"teamId" gorm:"foreignKey:team_id;OnDelete:Cascade;OnUpdate:CASCADE;references:team_id"`
}
