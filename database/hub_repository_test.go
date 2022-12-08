package database

import (
	"cubicasa/models"
	"testing"
)

func TestGetAllHub(t *testing.T) {
	hubRepo := HubRepository{}
	hubRepo.InitDB("PG")
	t.Run("Case1:(database is up)", func(t *testing.T) {
		_, isDone := hubRepo.GetAllHub()
		if isDone != true {
			t.Errorf("Get All hub(database connected) = %t ----> Result should be: %t", isDone, true)
		}
	})
	hubRepo.db = nil
	t.Run("Case2:(database is not connected)", func(t *testing.T) {
		_, isDone := hubRepo.GetAllHub()
		if isDone != false {
			t.Errorf("Get All hub(database disconnect) = %t ----> Result should be: %t", isDone, false)
		}
	})
}
func TestGetHubById(t *testing.T) {
	hubRepo := HubRepository{}
	hubRepo.InitDB("PG")
	t.Run("Case1: (id:365 not in database)", func(t *testing.T) {
		hub, isDone := hubRepo.GetHubById(365)
		if isDone != false {
			t.Errorf("get hub by id(%d)(not in data base) = %t ----> Result should be: %t", hub.HubID, isDone, false)
		}
	})
	t.Run("Case2: (id:2 in database)", func(t *testing.T) {
		hub, isDone := hubRepo.GetHubById(2)
		if isDone != true {
			t.Errorf("get hub by id(%d)(exits data base) = %t ----> Result should be: %t", hub.HubID, isDone, true)
		}
	})
	// nil db config for disconnect testcase
	hubRepo.db = nil
	t.Run("Case3: (database not connected)", func(t *testing.T) {
		hub, isDone := hubRepo.GetHubById(1)
		if isDone != false {
			t.Errorf("get hub by id(%d)(database not connect) = %t ----> Result should be: %t", hub.HubID, isDone, false)
		}
	})
}
func TestUpdateHub(t *testing.T) {
	hubRepo := HubRepository{}
	hubRepo.InitDB("PG")
	hub := models.Hub{
		HubName:   "Test Hub",
		HubRegion: "Some where",
		HubID:     4,
	}
	t.Run("Case1: (id:365 not in database)", func(t *testing.T) {
		id, isDone := hubRepo.UpdateHub(365, hub)
		if isDone != false {
			t.Errorf("Update hub(%d)(not in data base) = %t ----> Result should be: %t", id, isDone, false)
		}
	})

	t.Run("Case2: (id:1 in database and full info)", func(t *testing.T) {
		id, isDone := hubRepo.UpdateHub(1, hub)
		if isDone != true {
			t.Errorf("Update hub(%d)(exits data base) = %t ----> Result should be: %t", id, isDone, true)
		}
	})
	hub = models.Hub{
		HubRegion: "Some where miss",
	}
	t.Run("Case3: (id:1 in database but missed info)", func(t *testing.T) {
		id, isDone := hubRepo.UpdateHub(2, hub)
		if isDone != true {
			t.Errorf("Update hub(%d)(exits data base) = %t ----> Result should be: %t", id, isDone, true)
		}
	})
	// test generate new db config for disconnect testcase
	hubRepo.db = nil
	t.Run("Case4: (database not connected)", func(t *testing.T) {
		id, isDone := hubRepo.UpdateHub(1, hub)
		if isDone != false {
			t.Errorf("Update hub(%d)(database not connected) = %t ----> Result should be: %t", id, isDone, false)
		}
	})
}

func TestCreateHub(t *testing.T) {
	hubRepo := HubRepository{}
	hubRepo.InitDB("PG")
	hub := models.Hub{
		HubName:   "Test Hub",
		HubRegion: "Some where",
		HubID:     3,
	}
	t.Run("Case1: (database is connected)", func(t *testing.T) {
		hub, isDone := hubRepo.CreateHub(hub)
		if isDone != true {
			t.Errorf("create hub id: %d(database is connected) = %t ----> Result should be: %t", hub.HubID, isDone, true)
		}
	})
	hub = models.Hub{
		HubRegion: "missing hub name",
	}
	t.Run("Case2: (hub is miss info)", func(t *testing.T) {
		hub, isDone := hubRepo.CreateHub(hub)
		if isDone != false {
			t.Errorf("create hub id: (%d)(user miss info) = %t ----> Result should be: %t", hub.HubID, isDone, false)
		}
	})
	hubRepo.db = nil
	hub = models.Hub{
		HubName:   "Test Hub",
		HubRegion: "Some where",
		HubID:     4,
	}
	t.Run("Case3:(database is not connected)", func(t *testing.T) {
		hub, isDone := hubRepo.CreateHub(hub)
		if isDone != false {
			t.Errorf("create hub id: (%d)(database disconnected)= %t ----> Result should be: %t", hub.HubID, isDone, false)
		}
	})
}

func TestDeleteHub(t *testing.T) {
	hubRepo := HubRepository{}
	hubRepo.InitDB("PG")
	t.Run("Case 1: (id not in database)", func(t *testing.T) {
		id, isDone := hubRepo.DeleteHub(365)
		if isDone != false {
			t.Errorf("Delete hub(%d)(not in data base) = %t ----> Result should be: %t", id, isDone, false)
		}
	})
	t.Run("Case 2: (id in database)", func(t *testing.T) {
		id, isDone := hubRepo.DeleteHub(2)
		if isDone != true {
			t.Errorf("Delete hub(%d)(exits data base) = %t ----> Result should be: %t", id, isDone, true)
		}
	})
	hubRepo.db = nil
	t.Run("Case3:(database is not connected)", func(t *testing.T) {
		id, isDone := hubRepo.DeleteHub(1)
		if isDone != false {
			t.Errorf("Delete hub(%d)(database disconnect) = %t ----> Result should be: %t", id, isDone, false)
		}
	})
}
