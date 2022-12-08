package database

import (
	"cubicasa/models"
	"fmt"
	"gorm.io/gorm"
	"regexp"
	"strings"
)

// HubRepository hub create or update must provide new hub id
// --99 number marker failed
// can return a number marker(not only bool) for check what type of error raise

type HubRepository struct {
	db     *gorm.DB
	config Config
}

func (hubRepo *HubRepository) InitDB(dataBase string) {
	hubRepo.config = CreateConfig(dataBase)
	hubRepo.db = ConnectToDB(hubRepo.config)
}
func (hubRepo *HubRepository) GetAllHub() ([]models.Hub, bool) {
	var hubs []models.Hub
	if hubRepo.db == nil {
		return nil, false
	}
	result := hubRepo.db.Find(&hubs)
	if result.Error != nil {
		return nil, false
	} else {
		return hubs, true
	}
}
func (hubRepo *HubRepository) GetHubById(id int) (models.Hub, bool) {
	var hub models.Hub
	if hubRepo.db == nil {
		return hub, false
	}
	result := hubRepo.db.First(&hub, id)
	if result.Error != nil {
		return hub, false
	} else {
		return hub, true
	}
}
func (hubRepo *HubRepository) UpdateHub(id int, hub models.Hub) (int, bool) {
	if hub.HubID == 0 {
		hub.HubID = uint(id)
	}
	if hubRepo.db == nil || !ValidateId(id) || !validateHub(hub) {
		fmt.Println(hub.HubID)
		return -99, false
	}
	_, isExist := hubRepo.GetHubById(id)
	if isExist != true {
		return id, false
	}
	result := hubRepo.db.Where("hub_id = ?", id).Updates(&hub)
	if result.Error != nil {
		return id, false
	} else {
		return id, true
	}
}

func (hubRepo *HubRepository) CreateHub(hub models.Hub) (models.Hub, bool) {
	if hubRepo.db == nil || !validateHub(hub) {
		return hub, false
	}
	result := hubRepo.db.Create(&hub)
	if result.Error != nil {
		return hub, false
	} else {
		return hub, true
	}
}

func (hubRepo *HubRepository) DeleteHub(id int) (int, bool) {
	if hubRepo.db == nil || !ValidateId(id) {
		return id, false
	}
	_, isExist := hubRepo.GetHubById(id)
	if isExist != true {
		return id, false
	}
	// cascade on gorm not work
	result := hubRepo.db.Delete(&models.Hub{}, id)
	if result.Error != nil {
		return id, false
	} else {
		return id, true
	}
}

func validateHub(hub models.Hub) bool {
	if strings.Trim(hub.HubName, " ") == "" || len(hub.HubName) > 20 || !MatchRegex(hub.HubName) {
		return false
	}
	if strings.Trim(hub.HubRegion, " ") == "" || len(hub.HubRegion) > 20 || !matchRegexHubRegion(hub.HubRegion) {
		return false
	}
	return true
}
func matchRegexHubRegion(inputRegion string) bool {
	matched, _ := regexp.MatchString("[a-zA-Z]", inputRegion)
	return matched
}
