package service

import (
	"cubicasa/database"
	"cubicasa/models"
)

type TeamService struct {
	teamRepo database.TeamRepository
}

func (teamService *TeamService) InitService() {
	teamService.teamRepo.InitDB("PG")
}
func (teamService *TeamService) GetAllTeam() ([]models.Team, bool) {
	return teamService.teamRepo.GetAllTeam()
}
func (teamService *TeamService) GetTeamById(id int) (models.Team, bool) {
	return teamService.teamRepo.GetTeamById(id)
}

func (teamService *TeamService) UpdateTeam(id int, team models.Team) (int, bool) {
	return teamService.teamRepo.UpdateTeam(id, team)
}

func (teamService *TeamService) CreateTeam(team models.Team) (models.Team, bool) {
	return teamService.teamRepo.CreateTeam(team)
}

func (teamService *TeamService) DeleteTeam(id int) (int, bool) {
	return teamService.teamRepo.DeleteTeam(id)
}

func (teamService *TeamService) MoveTeam(teamId int, hubId int) bool {
	return teamService.teamRepo.MoveTeam(teamId, hubId)
}
