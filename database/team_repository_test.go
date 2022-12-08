package database

import (
	"cubicasa/models"
	"testing"
)

func TestGetAllTeam(t *testing.T) {
	teamRepo := TeamRepository{}
	teamRepo.InitDB("PG")
	t.Run("Case1:(database is up)", func(t *testing.T) {
		_, isDone := teamRepo.GetAllTeam()
		if isDone != true {
			t.Errorf("Get All team(database connected) = %t ----> Result should be: %t", isDone, true)
		}
	})
	teamRepo.db = nil
	t.Run("Case2:(database is not connected)", func(t *testing.T) {
		_, isDone := teamRepo.GetAllTeam()
		if isDone != false {
			t.Errorf("Get All team(database disconnect) = %t ----> Result should be: %t", isDone, false)
		}
	})
}

func TeamGetTeamById(t *testing.T) {
	teamRepo := TeamRepository{}
	teamRepo.InitDB("PG")
	t.Run("Case1: (id:365 not in database)", func(t *testing.T) {
		team, isDone := teamRepo.GetTeamById(365)
		if isDone != false {
			t.Errorf("get team by id(%d)(not in data base) = %t ----> Result should be: %t", team.TeamID, isDone, false)
		}
	})
	t.Run("Case2: (id:2 in database)", func(t *testing.T) {
		team, isDone := teamRepo.GetTeamById(2)
		if isDone != true {
			t.Errorf("get team by id(%d)(exits data base) = %t ----> Result should be: %t", team.TeamID, isDone, true)
		}
	})
	// nil db config for disconnect testcase
	teamRepo.db = nil
	t.Run("Case3: (database not connected)", func(t *testing.T) {
		team, isDone := teamRepo.GetTeamById(1)
		if isDone != false {
			t.Errorf("get team by id(%d)(database not connect) = %t ----> Result should be: %t", team.TeamID, isDone, false)
		}
	})
}

func TestUpdateTeam(t *testing.T) {
	teamRepo := TeamRepository{}
	teamRepo.InitDB("PG")
	team := models.Team{
		TeamName: "team test update",
		HubID:    1,
	}
	t.Run("Case1: (id:365 not in database)", func(t *testing.T) {
		id, isDone := teamRepo.UpdateTeam(365, team)
		if isDone != false {
			t.Errorf("Update team(%d)(not in data base) = %t ----> Result should be: %t", id, isDone, false)
		}
	})
	t.Run("Case2: (id:1 in database and full info)", func(t *testing.T) {
		id, isDone := teamRepo.UpdateTeam(1, team)
		if isDone != true {
			t.Errorf("Update team(%d)(exits data base) = %t ----> Result should be: %t", id, isDone, true)
		}
	})
	team = models.Team{
		TeamName: "test team update",
	}
	t.Run("Case3: (id:1 in database but missed info)", func(t *testing.T) {
		id, isDone := teamRepo.UpdateTeam(2, team)
		if isDone != true {
			t.Errorf("Update team(%d)(exits data base) = %t ----> Result should be: %t", id, isDone, true)
		}
	})
	// bind nil to disconnect db
	teamRepo.db = nil
	t.Run("Case4: (database not connected)", func(t *testing.T) {
		id, isDone := teamRepo.UpdateTeam(1, team)
		if isDone != false {
			t.Errorf("Update team(%d)(database not connected) = %t ----> Result should be: %t", id, isDone, false)
		}
	})
}

func TestCreateTeam(t *testing.T) {
	teamRepo := TeamRepository{}
	teamRepo.InitDB("PG")
	// hub 2 will  delete for test so create team will go to hub 1
	team := models.Team{
		TeamName: "team create test",
		HubID:    1,
	}
	t.Run("Case1: (database is connected)", func(t *testing.T) {
		team, isDone := teamRepo.CreateTeam(team)
		if isDone != true {
			t.Errorf("create team id: %d(database is connected) = %t ----> Result should be: %t", team.TeamID, isDone, true)
		}
	})
	team = models.Team{}
	t.Run("Case2: (team is miss info)", func(t *testing.T) {
		team, isDone := teamRepo.CreateTeam(team)
		if isDone != false {
			t.Errorf("create team id: (%d)(user miss info) = %t ----> Result should be: %t", team.TeamID, isDone, false)
		}
	})
	teamRepo.db = nil
	team = models.Team{
		TeamName: "team create test",
		HubID:    1,
	}
	t.Run("Case3:(database is not connected)", func(t *testing.T) {
		team, isDone := teamRepo.CreateTeam(team)
		if isDone != false {
			t.Errorf("create team id: (%d)(database disconnected)= %t ----> Result should be: %t", team.TeamID, isDone, false)
		}
	})
}

func TestDeleteTeam(t *testing.T) {
	teamRepo := TeamRepository{}
	teamRepo.InitDB("PG")
	t.Run("Case 1: (id not in database)", func(t *testing.T) {
		id, isDone := teamRepo.DeleteTeam(365)
		if isDone != false {
			t.Errorf("Delete team(%d)(not in data base) = %t ----> Result should be: %t", id, isDone, false)
		}
	})
	t.Run("Case 2: (id in database)", func(t *testing.T) {
		id, isDone := teamRepo.DeleteTeam(2)
		if isDone != true {
			t.Errorf("Delete team(%d)(exits data base) = %t ----> Result should be: %t", id, isDone, true)
		}
	})
	teamRepo.db = nil
	t.Run("Case3:(database is not connected)", func(t *testing.T) {
		id, isDone := teamRepo.DeleteTeam(5)
		if isDone != false {
			t.Errorf("Delete team(%d)(database disconnect) = %t ----> Result should be: %t", id, isDone, false)
		}
	})
}

func TestMoveTeam(t *testing.T) {
	teamRepo := TeamRepository{}
	teamRepo.InitDB("PG")
	teamId := 99
	hubId := 1
	t.Run("Case1: (hub not in database)", func(t *testing.T) {
		isDone := teamRepo.MoveTeam(teamId, hubId)
		if isDone != false {
			t.Errorf("move team(%d)(not in data base) to hub(%d)(not in database) = %t ----> Result should be: "+
				"%t", teamId, hubId, isDone, false)
		}
	})
	teamId = 2
	hubId = 99
	t.Run("Case2: (team id not in database)", func(t *testing.T) {
		isDone := teamRepo.MoveTeam(teamId, hubId)
		if isDone != false {
			t.Errorf("move team(%d)(not in data base) to hub(%d)(not in database) = %t ----> Result should be: "+
				"%t", teamId, hubId, isDone, false)
		}
	})
	teamId = 5
	hubId = 3
	t.Run("Case3: (hub id and team id in database)", func(t *testing.T) {
		isDone := teamRepo.MoveTeam(teamId, hubId)
		if isDone != true {
			t.Errorf("move team(%d)(not in data base) to hub(%d)(not in database) = %t ----> Result should be: "+
				"%t", teamId, hubId, isDone, false)
		}
	})

	teamRepo.db = nil
	teamId = 1
	hubId = 5
	t.Run("Case4:(database is not connected)", func(t *testing.T) {
		isDone := teamRepo.MoveTeam(teamId, hubId)
		if isDone != false {
			t.Errorf("move team(%d)(not in data base) to hub(%d)(not in database) = %t ----> Result should be: "+
				"%t", teamId, hubId, isDone, false)
		}
	})
}
