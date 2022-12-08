package database

import (
	"cubicasa/models"
	"testing"
)

func TestGetAllUser(t *testing.T) {
	userRepo := UserRepository{}
	userRepo.InitDB("PG")
	t.Run("Case1:(database is up)", func(t *testing.T) {
		_, isDone := userRepo.GetAllUser()
		if isDone != true {
			t.Errorf("Get All user(database connected) = %t ----> Result should be: %t", isDone, true)
		}
	})
	userRepo.db = nil
	t.Run("Case2:(database is not connected)", func(t *testing.T) {
		_, isDone := userRepo.GetAllUser()
		if isDone != false {
			t.Errorf("Get All user(database disconnect) = %t ----> Result should be: %t", isDone, false)
		}
	})
}
func TestGetUserById(t *testing.T) {
	userRepo := UserRepository{}
	userRepo.InitDB("PG")
	t.Run("Case1: (id:365 not in database)", func(t *testing.T) {
		user, isDone := userRepo.GetUserById(365)
		if isDone != false {
			t.Errorf("get user by id(%d)(not in data base) = %t ----> Result should be: %t", user.UserID, isDone, false)
		}
	})
	t.Run("Case2: (id:6 in database)", func(t *testing.T) {
		user, isDone := userRepo.GetUserById(3)
		if isDone != true {
			t.Errorf("get user by id(%d)(exits data base) = %t ----> Result should be: %t", user.UserID, isDone, true)
		}
	})
	// nil db config for disconnect testcase
	userRepo.db = nil
	t.Run("Case3: (database not connected)", func(t *testing.T) {
		user, isDone := userRepo.GetUserById(1)
		if isDone != false {
			t.Errorf("get user by id(%d)(database not connect) = %t ----> Result should be: %t", user.UserID, isDone, false)
		}
	})
}

func TestUpdateUser(t *testing.T) {
	userRepo := UserRepository{}
	userRepo.InitDB("PG")
	user := models.User{
		Role:     "test role",
		UserName: "test user",
		TeamId:   1,
	}
	t.Run("Case1: (id:365 not in database)", func(t *testing.T) {
		id, isDone := userRepo.UpdateUser(365, user)
		if isDone != false {
			t.Errorf("Update use(%d)(not in data base) = %t ----> Result should be: %t", id, isDone, false)
		}
	})
	t.Run("Case2: (id:1 in database and full info)", func(t *testing.T) {
		id, isDone := userRepo.UpdateUser(1, user)
		if isDone != true {
			t.Errorf("Update use(%d)(exits data base) = %t ----> Result should be: %t", id, isDone, true)
		}
	})
	user = models.User{
		UserName: "test user",
	}
	t.Run("Case3: (id:4 in database but missed info)", func(t *testing.T) {
		id, isDone := userRepo.UpdateUser(4, user)
		if isDone != true {
			t.Errorf("Update use(%d)(exits data base) = %t ----> Result should be: %t", id, isDone, true)
		}
	})
	// test generate new db config for disconnect testcase
	userRepo.db = nil
	t.Run("Case4: (database not connected)", func(t *testing.T) {
		id, isDone := userRepo.UpdateUser(1, user)
		if isDone != false {
			t.Errorf("Update use(%d)(database not connected) = %t ----> Result should be: %t", id, isDone, false)
		}
	})
}
func TestCreateUser(t *testing.T) {
	userRepo := UserRepository{}
	userRepo.InitDB("PG")
	user := models.User{
		Role:     "create role",
		UserName: "create user",
		TeamId:   1,
	}
	t.Run("Case1: (database is connected)", func(t *testing.T) {
		user, isDone := userRepo.CreateUser(user)
		if isDone != true {
			t.Errorf("create user id: %d(database is connected) = %t ----> Result should be: %t", user.UserID, isDone, true)
		}
	})
	user = models.User{
		Role: "missing role",
	}
	t.Run("Case2: (user is miss info)", func(t *testing.T) {
		user, isDone := userRepo.CreateUser(user)
		if isDone != false {
			t.Errorf("create user id: (%d)(user miss info) = %t ----> Result should be: %t", user.UserID, isDone, false)
		}
	})
	userRepo.db = nil
	user = models.User{
		Role:     "create role",
		UserName: "create user",
		TeamId:   2,
	}
	t.Run("Case3:(database is not connected)", func(t *testing.T) {
		user, isDone := userRepo.CreateUser(user)
		if isDone != false {
			t.Errorf("create user id: (%d)(database disconnected)= %t ----> Result should be: %t", user.UserID, isDone, false)
		}
	})
}

func TestDeleteUser(t *testing.T) {
	userRepo := UserRepository{}
	userRepo.InitDB("PG")
	t.Run("Case1: (id not in database)", func(t *testing.T) {
		id, isDone := userRepo.DeleteUser(365)
		if isDone != false {
			t.Errorf("Delete user(%d)(not in data base) = %t ----> Result should be: %t", id, isDone, false)
		}
	})
	t.Run("Case2: (id in database)", func(t *testing.T) {
		id, isDone := userRepo.DeleteUser(9)
		if isDone != true {
			t.Errorf("Delete user(%d)(exits data base) = %t ----> Result should be: %t", id, isDone, true)
		}
	})
	userRepo.db = nil
	t.Run("Case3:(database is not connected)", func(t *testing.T) {
		id, isDone := userRepo.DeleteUser(5)
		if isDone != false {
			t.Errorf("Delete user(%d)(database disconnect) = %t ----> Result should be: %t", id, isDone, false)
		}
	})
}

func TestMoveUser(t *testing.T) {
	userRepo := UserRepository{}
	userRepo.InitDB("PG")
	userId := 99
	teamId := 1
	t.Run("Case1: (userid not in database)", func(t *testing.T) {
		isDone := userRepo.MoveUser(userId, teamId)
		if isDone != false {
			t.Errorf("move user(%d)(not in data base) to team(%d)(not in database) = %t ----> Result should be: "+
				"%t", userId, teamId, isDone, false)
		}
	})
	userId = 2
	teamId = 99
	t.Run("Case2: (team id not in database)", func(t *testing.T) {
		isDone := userRepo.MoveUser(userId, teamId)
		if isDone != false {
			t.Errorf("move user(%d)(not in data base) to team(%d)(not in database) = %t ----> Result should be: "+
				"%t", userId, teamId, isDone, false)
		}
	})
	userId = 1
	teamId = 5
	t.Run("Case3: (user id and team id in database)", func(t *testing.T) {
		isDone := userRepo.MoveUser(userId, teamId)
		if isDone != true {
			t.Errorf("move user(%d)(not in data base) to team(%d)(not in database) = %t ----> Result should be: "+
				"%t", userId, teamId, isDone, false)
		}
	})

	userRepo.db = nil
	userId = 1
	teamId = 5
	t.Run("Case4:(database is not connected)", func(t *testing.T) {
		isDone := userRepo.MoveUser(userId, teamId)
		if isDone != false {
			t.Errorf("move user(%d)(not in data base) to team(%d)(not in database) = %t ----> Result should be: "+
				"%t", userId, teamId, isDone, false)
		}
	})
}
