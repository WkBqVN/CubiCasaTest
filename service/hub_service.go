package service

import (
	"cubicasa/database"
	"cubicasa/models"
)

type HubService struct {
	hubRepo database.HubRepository
}

func (hubService *HubService) InitService() {
	hubService.hubRepo.InitDB("PG")
}
func (hubService *HubService) GetAllHub() ([]models.Hub, bool) {
	return hubService.hubRepo.GetAllHub()
}

func (hubService *HubService) GetHubById(id int) (models.Hub, bool) {
	return hubService.hubRepo.GetHubById(id)
}
func (hubService *HubService) UpdateHub(id int, hub models.Hub) (int, bool) {
	return hubService.hubRepo.UpdateHub(id, hub)
}

func (hubService *HubService) CreateHub(hub models.Hub) (models.Hub, bool) {
	return hubService.hubRepo.CreateHub(hub)
}

func (hubService *HubService) DeleteHub(id int) (int, bool) {
	return hubService.hubRepo.DeleteHub(id)
}
