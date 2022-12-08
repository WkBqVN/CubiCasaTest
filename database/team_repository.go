package database

import (
	"cubicasa/models"
	"gorm.io/gorm"
	"strings"
)

// TeamRepository db must not nil , if nil all service will return false
// --99 number marker failed
// can return a number marker(not only bool) for check what type of error raise
// team create must provide hub id

type TeamRepository struct {
	db     *gorm.DB
	config Config
}

func (teamRepo *TeamRepository) InitDB(dataBase string) {
	teamRepo.config = CreateConfig(dataBase)
	teamRepo.db = ConnectToDB(teamRepo.config)
}
func (teamRepo *TeamRepository) GetAllTeam() ([]models.Team, bool) {
	var teams []models.Team
	if teamRepo.db == nil {
		return nil, false
	}
	result := teamRepo.db.Find(&teams)
	if result.Error != nil {
		return nil, false
	} else {
		return teams, true
	}
}

func (teamRepo *TeamRepository) GetTeamById(id int) (models.Team, bool) {
	var team models.Team
	if teamRepo.db == nil || !ValidateId(id) {
		return team, false
	}
	result := teamRepo.db.First(&team, id)
	if result.Error != nil {
		return team, false
	} else {
		return team, true
	}
}

func (teamRepo *TeamRepository) UpdateTeam(id int, team models.Team) (int, bool) {
	if ValidateId(id) {
		teamData, isExits := teamRepo.GetTeamById(id)
		if isExits != true {
			return id, false
		}
		if !validateTeam(team) {
			if team.HubID == 0 {
				team.HubID = teamData.HubID
			}
		}
	} else {
		return -99, false
	}
	if teamRepo.db == nil {
		return -99, false
	}
	result := teamRepo.db.Where("team_id = ?", id).Updates(&team)
	if result.Error != nil {
		return id, false
	} else {
		return id, true
	}
}

func (teamRepo *TeamRepository) CreateTeam(team models.Team) (models.Team, bool) {
	if teamRepo.db == nil || !validateTeam(team) {
		return team, false
	}
	result := teamRepo.db.Create(&team)
	if result.Error != nil {
		return team, false
	} else {
		return team, true
	}
}

func (teamRepo *TeamRepository) DeleteTeam(id int) (int, bool) {
	if teamRepo.db == nil || !ValidateId(id) {
		return -99, false
	}
	_, isExist := teamRepo.GetTeamById(id)
	if isExist != true {
		return id, false
	}
	result := teamRepo.db.Delete(&models.Team{}, id)
	if result.Error != nil {
		return id, false
	} else {
		return id, true
	}
}

func (teamRepo *TeamRepository) MoveTeam(teamId int, hubId int) bool {
	resultData, isDone := teamRepo.GetTeamById(teamId)
	if isDone == true {
		resultData.HubID = uint(hubId)
		_, isDone := teamRepo.UpdateTeam(teamId, resultData)
		if isDone == true {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func validateTeam(team models.Team) bool {
	// update bind on param path not put in json
	if team.TeamID != 0 {
		return false
	}
	if strings.Trim(team.TeamName, " ") == "" || len(team.TeamName) > 20 || !MatchRegex(team.TeamName) {
		return false
	}
	if team.HubID == 0 || team.HubID > 100 {
		return false
	}
	return true
}
